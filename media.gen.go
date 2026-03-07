package goshopee

import (
	"io"
)

type MediaService interface {
	// UploadImage Use this API to upload images and support different business scenarios through business and scene parameters.
	// Path: /api/v2/media/upload_image
	// https://open.shopee.com/documents/v2/v2.media.upload_image?module=130&type=1
	UploadImage(sid uint64, filename string, tok string) (*MediaUploadImageResponse, error)
	UploadImageFromReader(sid uint64, filename string, reader io.Reader, tok string) (*MediaUploadImageResponse, error)
	// InitVideoUpload Use this API to initialize a new video upload task and obtain a unique upload session ID.
	// Path: /api/v2/media/init_video_upload
	// https://open.shopee.com/documents/v2/v2.media.init_video_upload?module=130&type=1
	InitVideoUpload(sid uint64, filename string, tok string) (*MediaInitVideoUploadResponse, error)
	InitVideoUploadFromReader(sid uint64, filename string, reader io.Reader, tok string) (*MediaInitVideoUploadResponse, error)
	// UploadVideoPart Use this API to upload a video file in parts.
	// Path: /api/v2/media/upload_video_part
	// https://open.shopee.com/documents/v2/v2.media.upload_video_part?module=130&type=1
	UploadVideoPart(sid uint64, filename string, tok string) (*MediaUploadVideoPartResponse, error)
	UploadVideoPartFromReader(sid uint64, filename string, reader io.Reader, tok string) (*MediaUploadVideoPartResponse, error)
	// CompleteVideoUpload Use this API to finalize a video upload task. Once called, the system will transcode all uploaded video parts and prepare the video for use. All parts must be fully uploaded before calling this API.
	// Path: /api/v2/media/complete_video_upload
	// https://open.shopee.com/documents/v2/v2.media.complete_video_upload?module=130&type=1
	CompleteVideoUpload(sid uint64, filename string, tok string) (*MediaCompleteVideoUploadResponse, error)
	CompleteVideoUploadFromReader(sid uint64, filename string, reader io.Reader, tok string) (*MediaCompleteVideoUploadResponse, error)
	// GetVideoUploadResult Use this API to retrieve the current status and result of a video upload task. You can also use video_upload_result_push to get notified immediately when a video upload reaches a final status (SUCCEEDED, FAILED, or CANCELLED).
	// Path: /api/v2/media/get_video_upload_result
	// https://open.shopee.com/documents/v2/v2.media.get_video_upload_result?module=130&type=1
	GetVideoUploadResult(sid uint64, filename string, tok string) (*MediaGetVideoUploadResultResponse, error)
	GetVideoUploadResultFromReader(sid uint64, filename string, reader io.Reader, tok string) (*MediaGetVideoUploadResultResponse, error)
	// CancelVideoUpload Use this API to cancel an ongoing video upload task. If the upload has already succeeded, failed, or been cancelled, this operation will fail and return error.
	// Path: /api/v2/media/cancel_video_upload
	// https://open.shopee.com/documents/v2/v2.media.cancel_video_upload?module=130&type=1
	CancelVideoUpload(sid uint64, filename string, tok string) (*MediaCancelVideoUploadResponse, error)
	CancelVideoUploadFromReader(sid uint64, filename string, reader io.Reader, tok string) (*MediaCancelVideoUploadResponse, error)
}

type MediaServiceOp[T any] struct {
	client *Client[T]
}

// UploadImage Use this API to upload images and support different business scenarios through business and scene parameters.
// Path: /api/v2/media/upload_image
// https://open.shopee.com/documents/v2/v2.media.upload_image?module=130&type=1
func (s *MediaServiceOp[T]) UploadImage(sid uint64, filename string, tok string) (*MediaUploadImageResponse, error) {
	path := "/media/upload_image"
	resp := new(MediaUploadImageResponse)
	err := s.client.WithMerchant(sid, tok).Upload(path, "image", filename, resp)
	return resp, err
}

// UploadImageFromReader Use this API to upload images and support different business scenarios through business and scene parameters.
// Path: /api/v2/media/upload_image
// https://open.shopee.com/documents/v2/v2.media.upload_image?module=130&type=1
func (s *MediaServiceOp[T]) UploadImageFromReader(sid uint64, filename string, reader io.Reader, tok string) (*MediaUploadImageResponse, error) {
	path := "/media/upload_image"
	resp := new(MediaUploadImageResponse)
	err := s.client.WithMerchant(sid, tok).UploadFromReader(path, "image", filename, reader, resp)
	return resp, err
}

// InitVideoUpload Use this API to initialize a new video upload task and obtain a unique upload session ID.
// Path: /api/v2/media/init_video_upload
// https://open.shopee.com/documents/v2/v2.media.init_video_upload?module=130&type=1
func (s *MediaServiceOp[T]) InitVideoUpload(sid uint64, filename string, tok string) (*MediaInitVideoUploadResponse, error) {
	path := "/media/init_video_upload"
	resp := new(MediaInitVideoUploadResponse)
	err := s.client.WithMerchant(sid, tok).Upload(path, "image", filename, resp)
	return resp, err
}

// InitVideoUploadFromReader Use this API to initialize a new video upload task and obtain a unique upload session ID.
// Path: /api/v2/media/init_video_upload
// https://open.shopee.com/documents/v2/v2.media.init_video_upload?module=130&type=1
func (s *MediaServiceOp[T]) InitVideoUploadFromReader(sid uint64, filename string, reader io.Reader, tok string) (*MediaInitVideoUploadResponse, error) {
	path := "/media/init_video_upload"
	resp := new(MediaInitVideoUploadResponse)
	err := s.client.WithMerchant(sid, tok).UploadFromReader(path, "image", filename, reader, resp)
	return resp, err
}

// UploadVideoPart Use this API to upload a video file in parts.
// Path: /api/v2/media/upload_video_part
// https://open.shopee.com/documents/v2/v2.media.upload_video_part?module=130&type=1
func (s *MediaServiceOp[T]) UploadVideoPart(sid uint64, filename string, tok string) (*MediaUploadVideoPartResponse, error) {
	path := "/media/upload_video_part"
	resp := new(MediaUploadVideoPartResponse)
	err := s.client.WithMerchant(sid, tok).Upload(path, "image", filename, resp)
	return resp, err
}

// UploadVideoPartFromReader Use this API to upload a video file in parts.
// Path: /api/v2/media/upload_video_part
// https://open.shopee.com/documents/v2/v2.media.upload_video_part?module=130&type=1
func (s *MediaServiceOp[T]) UploadVideoPartFromReader(sid uint64, filename string, reader io.Reader, tok string) (*MediaUploadVideoPartResponse, error) {
	path := "/media/upload_video_part"
	resp := new(MediaUploadVideoPartResponse)
	err := s.client.WithMerchant(sid, tok).UploadFromReader(path, "image", filename, reader, resp)
	return resp, err
}

// CompleteVideoUpload Use this API to finalize a video upload task. Once called, the system will transcode all uploaded video parts and prepare the video for use. All parts must be fully uploaded before calling this API.
// Path: /api/v2/media/complete_video_upload
// https://open.shopee.com/documents/v2/v2.media.complete_video_upload?module=130&type=1
func (s *MediaServiceOp[T]) CompleteVideoUpload(sid uint64, filename string, tok string) (*MediaCompleteVideoUploadResponse, error) {
	path := "/media/complete_video_upload"
	resp := new(MediaCompleteVideoUploadResponse)
	err := s.client.WithMerchant(sid, tok).Upload(path, "image", filename, resp)
	return resp, err
}

// CompleteVideoUploadFromReader Use this API to finalize a video upload task. Once called, the system will transcode all uploaded video parts and prepare the video for use. All parts must be fully uploaded before calling this API.
// Path: /api/v2/media/complete_video_upload
// https://open.shopee.com/documents/v2/v2.media.complete_video_upload?module=130&type=1
func (s *MediaServiceOp[T]) CompleteVideoUploadFromReader(sid uint64, filename string, reader io.Reader, tok string) (*MediaCompleteVideoUploadResponse, error) {
	path := "/media/complete_video_upload"
	resp := new(MediaCompleteVideoUploadResponse)
	err := s.client.WithMerchant(sid, tok).UploadFromReader(path, "image", filename, reader, resp)
	return resp, err
}

// GetVideoUploadResult Use this API to retrieve the current status and result of a video upload task. You can also use video_upload_result_push to get notified immediately when a video upload reaches a final status (SUCCEEDED, FAILED, or CANCELLED).
// Path: /api/v2/media/get_video_upload_result
// https://open.shopee.com/documents/v2/v2.media.get_video_upload_result?module=130&type=1
func (s *MediaServiceOp[T]) GetVideoUploadResult(sid uint64, filename string, tok string) (*MediaGetVideoUploadResultResponse, error) {
	path := "/media/get_video_upload_result"
	resp := new(MediaGetVideoUploadResultResponse)
	err := s.client.WithMerchant(sid, tok).Upload(path, "image", filename, resp)
	return resp, err
}

// GetVideoUploadResultFromReader Use this API to retrieve the current status and result of a video upload task. You can also use video_upload_result_push to get notified immediately when a video upload reaches a final status (SUCCEEDED, FAILED, or CANCELLED).
// Path: /api/v2/media/get_video_upload_result
// https://open.shopee.com/documents/v2/v2.media.get_video_upload_result?module=130&type=1
func (s *MediaServiceOp[T]) GetVideoUploadResultFromReader(sid uint64, filename string, reader io.Reader, tok string) (*MediaGetVideoUploadResultResponse, error) {
	path := "/media/get_video_upload_result"
	resp := new(MediaGetVideoUploadResultResponse)
	err := s.client.WithMerchant(sid, tok).UploadFromReader(path, "image", filename, reader, resp)
	return resp, err
}

// CancelVideoUpload Use this API to cancel an ongoing video upload task. If the upload has already succeeded, failed, or been cancelled, this operation will fail and return error.
// Path: /api/v2/media/cancel_video_upload
// https://open.shopee.com/documents/v2/v2.media.cancel_video_upload?module=130&type=1
func (s *MediaServiceOp[T]) CancelVideoUpload(sid uint64, filename string, tok string) (*MediaCancelVideoUploadResponse, error) {
	path := "/media/cancel_video_upload"
	resp := new(MediaCancelVideoUploadResponse)
	err := s.client.WithMerchant(sid, tok).Upload(path, "image", filename, resp)
	return resp, err
}

// CancelVideoUploadFromReader Use this API to cancel an ongoing video upload task. If the upload has already succeeded, failed, or been cancelled, this operation will fail and return error.
// Path: /api/v2/media/cancel_video_upload
// https://open.shopee.com/documents/v2/v2.media.cancel_video_upload?module=130&type=1
func (s *MediaServiceOp[T]) CancelVideoUploadFromReader(sid uint64, filename string, reader io.Reader, tok string) (*MediaCancelVideoUploadResponse, error) {
	path := "/media/cancel_video_upload"
	resp := new(MediaCancelVideoUploadResponse)
	err := s.client.WithMerchant(sid, tok).UploadFromReader(path, "image", filename, reader, resp)
	return resp, err
}
