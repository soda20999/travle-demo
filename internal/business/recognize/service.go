package recognize


import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"time"

	postgresql "iam/internal/pkg/config/postsql"
	dicover "iam/internal/business/discover"
	"iam/pkg/config"

	"gorm.io/gorm"
)

var recognizerClient = &http.Client{Timeout: 30 * time.Second}

func CreateAttractionImage(img *AttractionImage) error {
	if err := postgresql.DB.Create(img).Error; err != nil {
		return err
	}
	if err := indexAdd(img.ID, img.ImageURL); err != nil {
		postgresql.DB.Delete(img)
		return fmt.Errorf("index add failed: %w", err)
	}
	return nil
}

func DeleteAttractionImage(id int64) error {
	var img AttractionImage
	if err := postgresql.DB.First(&img, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return err
	}
	if err := indexRemove(id); err != nil {
		return fmt.Errorf("index remove failed: %w", err)
	}
	return postgresql.DB.Delete(&img).Error
}

func GetImagesByAttractionID(attractionID int64) ([]AttractionImage, error) {
	var images []AttractionImage
	err := postgresql.DB.Where("attraction_id = ?", attractionID).
		Order("created_at DESC").Find(&images).Error
	return images, err
}

func RebuildIndex() error {
	baseURL := config.Conf.RecognizerConfig.BaseURL
	resp, err := recognizerClient.Post(baseURL+"/index/rebuild", "application/json", nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("rebuild failed: %s", body)
	}
	return nil
}

type PredictResult struct {
	ImageID    int64   `json:"image_id"`
	Similarity float64 `json:"similarity"`
}

type RecognizeResult struct {
	Attraction dicover.Attraction `json:"attraction"`
	Similarity float64                   `json:"similarity"`
}

func Recognize(fileHeader *multipart.FileHeader, topK int) ([]RecognizeResult, error) {
	predictResults, err := callPredict(fileHeader, topK)
	if err != nil {
		return nil, err
	}
	if len(predictResults) == 0 {
		return []RecognizeResult{}, nil
	}

	imageIDs := make([]int64, len(predictResults))
	for i, r := range predictResults {
		imageIDs[i] = r.ImageID
	}

	var images []AttractionImage
	if err := postgresql.DB.Where("id IN ?", imageIDs).Find(&images).Error; err != nil {
		return nil, err
	}

	attractionIDs := make([]int64, 0, len(images))
	imageToAttraction := make(map[int64]int64)
	for _, img := range images {
		imageToAttraction[img.ID] = img.AttractionID
		attractionIDs = append(attractionIDs, img.AttractionID)
	}

	var attractions []dicover.Attraction
	if err := postgresql.DB.Where("id IN ?", attractionIDs).Find(&attractions).Error; err != nil {
		return nil, err
	}
	attractionMap := make(map[int64]dicover.Attraction)
	for _, a := range attractions {
		attractionMap[a.ID] = a
	}

	seen := make(map[int64]bool)
	results := make([]RecognizeResult, 0, len(predictResults))
	for _, pr := range predictResults {
		attrID, ok := imageToAttraction[pr.ImageID]
		if !ok {
			continue
		}
		if seen[attrID] {
			continue
		}
		attr, ok := attractionMap[attrID]
		if !ok {
			continue
		}
		seen[attrID] = true
		results = append(results, RecognizeResult{
			Attraction: attr,
			Similarity: pr.Similarity,
		})
	}
	return results, nil
}

func callPredict(fileHeader *multipart.FileHeader, topK int) ([]PredictResult, error) {
	baseURL := config.Conf.RecognizerConfig.BaseURL

	file, err := fileHeader.Open()
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)
	part, err := writer.CreateFormFile("image", fileHeader.Filename)
	if err != nil {
		return nil, err
	}
	if _, err := io.Copy(part, file); err != nil {
		return nil, err
	}
	if err := writer.WriteField("top_k", fmt.Sprintf("%d", topK)); err != nil {
		return nil, err
	}
	if err := writer.Close(); err != nil {
		return nil, err
	}

	resp, err := recognizerClient.Post(baseURL+"/predict", writer.FormDataContentType(), &buf)
	if err != nil {
		return nil, fmt.Errorf("recognizer service unavailable: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("predict failed (%d): %s", resp.StatusCode, body)
	}

	var result struct {
		Results []PredictResult `json:"results"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return result.Results, nil
}

func indexAdd(imageID int64, imageURL string) error {
	baseURL := config.Conf.RecognizerConfig.BaseURL
	body, _ := json.Marshal(map[string]interface{}{
		"image_id":  imageID,
		"image_url": imageURL,
	})
	resp, err := recognizerClient.Post(baseURL+"/index/add", "application/json", bytes.NewReader(body))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		respBody, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("index add failed: %s", respBody)
	}
	return nil
}

func indexRemove(imageID int64) error {
	baseURL := config.Conf.RecognizerConfig.BaseURL
	body, _ := json.Marshal(map[string]interface{}{
		"image_id": imageID,
	})
	resp, err := recognizerClient.Post(baseURL+"/index/remove", "application/json", bytes.NewReader(body))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		respBody, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("index remove failed: %s", respBody)
	}
	return nil
}
