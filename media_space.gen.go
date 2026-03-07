package goshopee

import (
	"io"
)

type MediaSpaceService interface {
	// InitVideoUpload Initiate video upload session.
	//
	// Video duration should be between 10s and 60s (inclusive).
	// Path: /api/v2/media_space/init_video_upload
	// https://open.shopee.com/documents/v2/v2.media_space.init_video_upload?module=91&type=1
	InitVideoUpload(sid uint64, filename string, tok string) (*InitVideoUploadResponse, error)
	InitVideoUploadFromReader(sid uint64, filename string, reader io.Reader, tok string) (*InitVideoUploadResponse, error)
	// UploadVideoPart Upload video file by part using the upload_id in initiate_video_upload.
	//
	// The request Content-Type of this API should be of multipart/form-data
	//
	//
	// Path: /api/v2/media_space/upload_video_part
	// https://open.shopee.com/documents/v2/v2.media_space.upload_video_part?module=91&type=1
	UploadVideoPart(sid uint64, filename string, tok string) (*UploadVideoPartResponse, error)
	UploadVideoPartFromReader(sid uint64, filename string, reader io.Reader, tok string) (*UploadVideoPartResponse, error)
	// CompleteVideoUpload Complete the video upload and starts the transcoding process when all parts are uploaded successfully.
	// Path: /api/v2/media_space/complete_video_upload
	// https://open.shopee.com/documents/v2/v2.media_space.complete_video_upload?module=91&type=1
	CompleteVideoUpload(sid uint64, filename string, tok string) (*CompleteVideoUploadResponse, error)
	CompleteVideoUploadFromReader(sid uint64, filename string, reader io.Reader, tok string) (*CompleteVideoUploadResponse, error)
	// GetVideoUploadResult Query the upload status and result of video upload.
	// Path: /api/v2/media_space/get_video_upload_result
	// https://open.shopee.com/documents/v2/v2.media_space.get_video_upload_result?module=91&type=1
	GetVideoUploadResult(sid uint64, filename string, tok string) (*GetVideoUploadResultResponse, error)
	GetVideoUploadResultFromReader(sid uint64, filename string, reader io.Reader, tok string) (*GetVideoUploadResultResponse, error)
	// CancelVideoUpload Cancel a video upload session
	// Path: /api/v2/media_space/cancel_video_upload
	// https://open.shopee.com/documents/v2/v2.media_space.cancel_video_upload?module=91&type=1
	CancelVideoUpload(sid uint64, filename string, tok string) (*CancelVideoUploadResponse, error)
	CancelVideoUploadFromReader(sid uint64, filename string, reader io.Reader, tok string) (*CancelVideoUploadResponse, error)
	// UploadImage Use this API to upload multiple image files (less than 9 images).
	// Path: /api/v2/media_space/upload_image
	// https://open.shopee.com/documents/v2/v2.media_space.upload_image?module=91&type=1
	UploadImage(sid uint64, filename string, tok string) (*SharedUploadImageResponse, error)
	UploadImageFromReader(sid uint64, filename string, reader io.Reader, tok string) (*SharedUploadImageResponse, error)
}

type MediaSpaceServiceOp[T any] struct {
	client *Client[T]
}

// InitVideoUpload Initiate video upload session.
//
// Video duration should be between 10s and 60s (inclusive).
// Path: /api/v2/media_space/init_video_upload
// https://open.shopee.com/documents/v2/v2.media_space.init_video_upload?module=91&type=1
func (s *MediaSpaceServiceOp[T]) InitVideoUpload(sid uint64, filename string, tok string) (*InitVideoUploadResponse, error) {
	path := "/media_space/init_video_upload"
	resp := new(InitVideoUploadResponse)
	err := s.client.WithShop(sid, tok).Upload(path, "image", filename, resp)
	return resp, err
}

// InitVideoUploadFromReader Initiate video upload session.
//
// Video duration should be between 10s and 60s (inclusive).
// Path: /api/v2/media_space/init_video_upload
// https://open.shopee.com/documents/v2/v2.media_space.init_video_upload?module=91&type=1
func (s *MediaSpaceServiceOp[T]) InitVideoUploadFromReader(sid uint64, filename string, reader io.Reader, tok string) (*InitVideoUploadResponse, error) {
	path := "/media_space/init_video_upload"
	resp := new(InitVideoUploadResponse)
	err := s.client.WithShop(sid, tok).UploadFromReader(path, "image", filename, reader, resp)
	return resp, err
}

// UploadVideoPart Upload video file by part using the upload_id in initiate_video_upload.
//
// The request Content-Type of this API should be of multipart/form-data
//
// Path: /api/v2/media_space/upload_video_part
// https://open.shopee.com/documents/v2/v2.media_space.upload_video_part?module=91&type=1
func (s *MediaSpaceServiceOp[T]) UploadVideoPart(sid uint64, filename string, tok string) (*UploadVideoPartResponse, error) {
	path := "/media_space/upload_video_part"
	resp := new(UploadVideoPartResponse)
	err := s.client.WithMerchant(sid, tok).Upload(path, "image", filename, resp)
	return resp, err
}

// UploadVideoPartFromReader Upload video file by part using the upload_id in initiate_video_upload.
//
// The request Content-Type of this API should be of multipart/form-data
//
// Path: /api/v2/media_space/upload_video_part
// https://open.shopee.com/documents/v2/v2.media_space.upload_video_part?module=91&type=1
func (s *MediaSpaceServiceOp[T]) UploadVideoPartFromReader(sid uint64, filename string, reader io.Reader, tok string) (*UploadVideoPartResponse, error) {
	path := "/media_space/upload_video_part"
	resp := new(UploadVideoPartResponse)
	err := s.client.WithMerchant(sid, tok).UploadFromReader(path, "image", filename, reader, resp)
	return resp, err
}

// CompleteVideoUpload Complete the video upload and starts the transcoding process when all parts are uploaded successfully.
// Path: /api/v2/media_space/complete_video_upload
// https://open.shopee.com/documents/v2/v2.media_space.complete_video_upload?module=91&type=1
func (s *MediaSpaceServiceOp[T]) CompleteVideoUpload(sid uint64, filename string, tok string) (*CompleteVideoUploadResponse, error) {
	path := "/media_space/complete_video_upload"
	resp := new(CompleteVideoUploadResponse)
	err := s.client.WithMerchant(sid, tok).Upload(path, "image", filename, resp)
	return resp, err
}

// CompleteVideoUploadFromReader Complete the video upload and starts the transcoding process when all parts are uploaded successfully.
// Path: /api/v2/media_space/complete_video_upload
// https://open.shopee.com/documents/v2/v2.media_space.complete_video_upload?module=91&type=1
func (s *MediaSpaceServiceOp[T]) CompleteVideoUploadFromReader(sid uint64, filename string, reader io.Reader, tok string) (*CompleteVideoUploadResponse, error) {
	path := "/media_space/complete_video_upload"
	resp := new(CompleteVideoUploadResponse)
	err := s.client.WithMerchant(sid, tok).UploadFromReader(path, "image", filename, reader, resp)
	return resp, err
}

// GetVideoUploadResult Query the upload status and result of video upload.
// Path: /api/v2/media_space/get_video_upload_result
// https://open.shopee.com/documents/v2/v2.media_space.get_video_upload_result?module=91&type=1
func (s *MediaSpaceServiceOp[T]) GetVideoUploadResult(sid uint64, filename string, tok string) (*GetVideoUploadResultResponse, error) {
	path := "/media_space/get_video_upload_result"
	resp := new(GetVideoUploadResultResponse)
	err := s.client.WithShop(sid, tok).Upload(path, "image", filename, resp)
	return resp, err
}

// GetVideoUploadResultFromReader Query the upload status and result of video upload.
// Path: /api/v2/media_space/get_video_upload_result
// https://open.shopee.com/documents/v2/v2.media_space.get_video_upload_result?module=91&type=1
func (s *MediaSpaceServiceOp[T]) GetVideoUploadResultFromReader(sid uint64, filename string, reader io.Reader, tok string) (*GetVideoUploadResultResponse, error) {
	path := "/media_space/get_video_upload_result"
	resp := new(GetVideoUploadResultResponse)
	err := s.client.WithShop(sid, tok).UploadFromReader(path, "image", filename, reader, resp)
	return resp, err
}

// CancelVideoUpload Cancel a video upload session
// Path: /api/v2/media_space/cancel_video_upload
// https://open.shopee.com/documents/v2/v2.media_space.cancel_video_upload?module=91&type=1
func (s *MediaSpaceServiceOp[T]) CancelVideoUpload(sid uint64, filename string, tok string) (*CancelVideoUploadResponse, error) {
	path := "/media_space/cancel_video_upload"
	resp := new(CancelVideoUploadResponse)
	err := s.client.WithShop(sid, tok).Upload(path, "image", filename, resp)
	return resp, err
}

// CancelVideoUploadFromReader Cancel a video upload session
// Path: /api/v2/media_space/cancel_video_upload
// https://open.shopee.com/documents/v2/v2.media_space.cancel_video_upload?module=91&type=1
func (s *MediaSpaceServiceOp[T]) CancelVideoUploadFromReader(sid uint64, filename string, reader io.Reader, tok string) (*CancelVideoUploadResponse, error) {
	path := "/media_space/cancel_video_upload"
	resp := new(CancelVideoUploadResponse)
	err := s.client.WithShop(sid, tok).UploadFromReader(path, "image", filename, reader, resp)
	return resp, err
}

// UploadImage Use this API to upload multiple image files (less than 9 images).
// Path: /api/v2/media_space/upload_image
// https://open.shopee.com/documents/v2/v2.media_space.upload_image?module=91&type=1
func (s *MediaSpaceServiceOp[T]) UploadImage(sid uint64, filename string, tok string) (*SharedUploadImageResponse, error) {
	path := "/media_space/upload_image"
	resp := new(SharedUploadImageResponse)
	err := s.client.WithMerchant(sid, tok).Upload(path, "image", filename, resp)
	return resp, err
}

// UploadImageFromReader Use this API to upload multiple image files (less than 9 images).
// Path: /api/v2/media_space/upload_image
// https://open.shopee.com/documents/v2/v2.media_space.upload_image?module=91&type=1
func (s *MediaSpaceServiceOp[T]) UploadImageFromReader(sid uint64, filename string, reader io.Reader, tok string) (*SharedUploadImageResponse, error) {
	path := "/media_space/upload_image"
	resp := new(SharedUploadImageResponse)
	err := s.client.WithMerchant(sid, tok).UploadFromReader(path, "image", filename, reader, resp)
	return resp, err
}
