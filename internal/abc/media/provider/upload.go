package media_provider

import (
	"context"
	"github.com/semenovem/portal/internal/abc/media"
)

func (p *MediaProvider) CreateUploadedFile(
	ctx context.Context,
	s3Path, note string,
	mediaObjType media.MediaObjectType,
) (uploadedFileID uint32, err error) {
	sq := `INSERT INTO media.upload_files (note, kind, s3_path)
		VALUES ($1, $2, $3) returning id;`

	if err = p.db.QueryRow(ctx, sq, note, mediaObjType, s3Path).Scan(&uploadedFileID); err != nil {
		p.logger.Named("CreateUploadedFile").DB(err)
		return 0, err
	}

	return
}
