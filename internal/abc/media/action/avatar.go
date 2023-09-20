package media_action

import (
	"context"
	"fmt"
	"github.com/semenovem/portal/internal/abc/media"
	"io"
)

func (a *MediaAction) UploadAvatar(
	ctx context.Context,
	thisUserID uint32,
	mediaObj media.ObjectType,
	file io.Reader,
) (avatarID uint32, previewContentBase64 string, err error) {
	var (
		ll = a.logger.Func(ctx, "UploadAvatar").With("mediaObj", mediaObj)
	)

	fmt.Println(">>>>>>>>>> ", mediaObj)

	//img, _, err := image.Decode(file)
	//if err != nil {
	//	ll.Error(err.Error())
	//	return 0, "", err
	//}
	//
	//size := img.Bounds().Size()
	//if size.Y < minWidthPx || size.Y < minHeightPx {
	//	return nil, failure.NewError(http.StatusBadRequest,
	//		failure.ImageDimensionsLessMinResponse,
	//		failure.Args{minWidthPx, minHeightPx})
	//}
	//
	//newImage := resize.Thumbnail(ct.config.Avatar.MaxWidthPixels, ct.config.Avatar.MaxHeightPixels, img, resize.Lanczos3)
	//
	//

	ll.Named("tmp")

	return 0, "", err
}
