package media_provider

import (
	"context"
	"github.com/semenovem/portal/internal/abc/media"
)

func (p *MediaProvider) InsertAvatar(ctx context.Context, avatarID string) error {
	sq := `INSERT INTO media.avatars (id) VALUES ($1) ON CONFLICT DO NOTHING;`

	if _, err := p.db.Exec(ctx, sq, avatarID); err != nil {
		p.logger.Func(ctx, "InsertAvatar").DB(err)
		return err
	}

	return nil
}

func (p *MediaProvider) CreateUploadedFile(
	ctx context.Context,
	s3Path, note string,
	mediaObjType media.ObjectType,
) (uploadedFileID uint32, err error) {
	sq := `INSERT INTO media.upload_files (note, kind, s3_path)
		VALUES ($1, $2, $3) returning id;`

	if err = p.db.QueryRow(ctx, sq, note, mediaObjType, s3Path).Scan(&uploadedFileID); err != nil {
		p.logger.Func(ctx, "CreateUploadedFile").DB(err)
		return 0, err
	}

	return
}
