package srv

import (
	"context"
	"github.com/iissy/goweb/src/domain"
	"github.com/iissy/goweb/src/model"
	"github.com/iissy/goweb/src/redis"
)

func (s *Hrefs) GetToken(ctx context.Context, req string, rsp *string) error {
	result, err := redis.Get(req)
	*rsp = result
	return err
}

func (s *Hrefs) SetToken(ctx context.Context, req *model.Token, rsp *bool) error {
	err := redis.Set(req.UserId, req.Code)
	return err
}

func (s *Hrefs) GetAccount(ctx context.Context, req int, rsp *model.Account) error {
	result, err := domain.GetAccount(req)
	*rsp = *result
	return err
}

func (s *Hrefs) GetAccountList(ctx context.Context, req *model.Pager, rsp *model.AccountList) error {
	result, err := domain.GetAccountList(req.Page, req.Size)
	rsp.List = result.List
	rsp.Total = result.Total
	return err
}

func (s *Hrefs) GetArticleList(ctx context.Context, req *model.SearchPager, rsp *model.ArticleList) error {
	result, err := domain.GetArticleList(req.Page, req.Size, req.Search)
	rsp.List = result.List
	rsp.Total = result.Total
	return err
}

func (s *Hrefs) DeleteArticle(ctx context.Context, req string, rsp *int64) error {
	result, err := domain.DeleteArticle(req)
	*rsp = result
	return err
}

func (s *Hrefs) SaveArticle(ctx context.Context, req *model.Article, rsp *int64) error {
	result, err := domain.SaveArticle(req)
	*rsp = result
	return err
}

func (s *Hrefs) Login(ctx context.Context, req *model.Account, rsp *model.Account) error {
	result, err := domain.Login(req)
	*rsp = *result
	return err
}

func (s *Hrefs) GetCusLink(ctx context.Context, req int, rsp *model.CusLink) error {
	result, err := domain.GetCusLink(req)
	*rsp = *result
	return err
}

func (s *Hrefs) GetCusLinkList(ctx context.Context, req *model.SearchPager, rsp *model.CusLinkList) error {
	result, err := domain.GetCusLinkList(req.Page, req.Size, req.Search)
	rsp.List = result.List
	rsp.Total = result.Total
	return err
}

func (s *Hrefs) DeleteCusLink(ctx context.Context, req int, rsp *int64) error {
	result, err := domain.DeleteCusLink(req)
	*rsp = result
	return err
}

func (s *Hrefs) SaveCusLink(ctx context.Context, req *model.CusLink, rsp *int64) error {
	result, err := domain.SaveCusLink(req)
	*rsp = result
	return err
}

func (s *Hrefs) GetLinkCat(ctx context.Context, req string, rsp *model.LinkCat) error {
	result, err := domain.GetLinkCat(req)
	*rsp = *result
	return err
}

func (s *Hrefs) GetLink(ctx context.Context, req string, rsp *model.Link) error {
	result, err := domain.GetLink(req)
	*rsp = *result
	return err
}

func (s *Hrefs) GetLinkList(ctx context.Context, req *model.SearchPager, rsp *model.LinkList) error {
	result, err := domain.GetLinkList(req.Page, req.Size, req.Search)
	rsp.List = result.List
	rsp.Total = result.Total
	return err
}

func (s *Hrefs) DeleteLink(ctx context.Context, req string, rsp *int64) error {
	result, err := domain.DeleteLink(req)
	*rsp = result
	return err
}

func (s *Hrefs) SaveLink(ctx context.Context, req *model.Link, rsp *int64) error {
	result, err := domain.SaveLink(req)
	*rsp = result
	return err
}

func (s *Hrefs) GetCatOptions(ctx context.Context, req bool, rsp *model.LinkCatList) error {
	items, err := domain.GetCatOptions()
	if err != nil {
		return err
	}

	rsp.List = items

	return nil
}
