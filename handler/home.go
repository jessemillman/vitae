package handler

import (
	"encoding/json"
	"os"

	"github.com/jessemillman/vitae/types"
	"github.com/jessemillman/vitae/view/home"
	"github.com/labstack/echo/v4"
)

type HomeHandler struct{}

func (h HomeHandler) HandleHomeShow(c echo.Context) error {
	jsonData, err := os.ReadFile("profile.json")
	if err != nil {
		return err
	}

	var cv types.CV

	err = json.Unmarshal(jsonData, &cv)
	if err != nil {
		return err
	}

	return render(c, home.Show(cv))
}
