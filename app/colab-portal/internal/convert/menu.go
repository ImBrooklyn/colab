package convert

import (
    "colab/api/grpc/portal"
    "colab/app/pkg/data/model"
)

func MenuModelToMenuItem(menuModel *model.Menu) *portal.UserMenuItem {
    return &portal.UserMenuItem{
        Id:            menuModel.ID,
        Pid:           menuModel.Pid,
        Title:         menuModel.Title,
        Icon:          menuModel.Icon,
        Url:           menuModel.URL,
        Filepath:      menuModel.Filepath,
        Params:        menuModel.Params,
        Node:          *menuModel.Node,
        Sort:          menuModel.Sort,
        Status:        *menuModel.Status,
        IsInner:       menuModel.IsInner,
        DefaultValues: menuModel.DefaultValues,
        ShowSlider:    *menuModel.ShowSlider,
        StatusText:    getStatusText(*menuModel.Status),
        InnerText:     getInnerText(menuModel.IsInner),
        FullUrl:       getFullUrl(menuModel.URL, menuModel.Params, menuModel.DefaultValues),
        Children:      make([]*portal.UserMenuItem, 0),
    }
}

func getFullUrl(url string, params string, values string) string {
    if (params != "" && values != "") || values != "" {
        return url + "/" + values
    }
    return url
}

func getInnerText(inner bool) string {
    if inner {
        return "内页"
    } else {
        return "导航"
    }
}

func getStatusText(status int32) string {
    switch status {
    case 0:
        return "禁用"
    case 1:
        return "使用中"
    default:
        return ""
    }
}

func MenuModelsToMenuItems(menuModels []*model.Menu) []*portal.UserMenuItem {
    menuMap := make(map[int64]*portal.UserMenuItem, len(menuModels))
    var menuList []*portal.UserMenuItem
    for _, item := range menuModels {
        menuItem := MenuModelToMenuItem(item)
        menuMap[item.ID] = menuItem
        menuList = append(menuList, menuItem)
    }

    list := make([]*portal.UserMenuItem, 0)
    for _, item := range menuList {
        if item.Pid == 0 {
            list = append(list, item)
            continue
        }

        if parentMenu, ok := menuMap[item.Pid]; ok {
            parentMenu.Children = append(parentMenu.Children, item)
        }
    }

    return list
}
