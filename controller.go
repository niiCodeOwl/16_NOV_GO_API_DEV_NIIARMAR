package controller

// func generateId() int {
// 	r := rand.New(rand.NewSource(time.Now().UnixNano()))
// 	return r.Intn(10000)
//   }
  
//   func getPlayers(c echo.Context) error {
// 	return c.JSON(http.StatusOK, players)
//   }
  
//   func postPlayer(c echo.Context) error {
// 	player := Player{}
// 	err := c.Bind(&player)
// 	if err != nil {
// 	  return echo.NewHTTPError(http.StatusUnprocessableEntity)
// 	}
// 	player.Id = generateId()
// 	players = append(players, player)
// 	return c.JSON(http.StatusCreated, players)
//   }