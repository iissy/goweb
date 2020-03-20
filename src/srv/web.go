package srv

import (
	"context"
	"hrefs.cn/src/domain"
	"hrefs.cn/src/model"
)

func (s *Hrefs) IndexLinks(ctx context.Context, req interface{}, rsp *model.LinkItems) error {
	items, err := domain.IndexLinks()
	if err != nil {
		return err
	}

	rsp.Items = items

	return nil
}

func (s *Hrefs) TopArticles(ctx context.Context, req interface{}, rsp *model.ArticleItems) error {
	items, err := domain.TopArticles()
	if err != nil {
		return err
	}

	rsp.Items = items

	return nil
}

func (s *Hrefs) TopCusLinks(ctx context.Context, req interface{}, rsp *model.CusLinkItems) error {
	items, err := domain.TopCusLinks()
	if err != nil {
		return err
	}

	rsp.Items = items

	return nil
}
