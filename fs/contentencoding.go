package fs

import "context"

func ContentEncoding(ctx context.Context, o ObjectInfo) string {
	if ce, ok := o.(ContentEncodinger); ok {
		return ce.ContentEncoding(ctx)
	}
	return ""
}
