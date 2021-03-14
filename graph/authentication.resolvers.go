package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/go-gomail/gomail"
	"github.com/sgabriella27/TPAWebGA_Back/database"
	"github.com/sgabriella27/TPAWebGA_Back/graph/generated"
	"github.com/sgabriella27/TPAWebGA_Back/graph/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func (r *communityAssetResolver) User(ctx context.Context, obj *model.CommunityAsset) (*model.User, error) {
	user := new(model.User)

	return user, database.GetDatabase().First(user, obj.UserID).Error
}

func (r *communityAssetResolver) Comments(ctx context.Context, obj *model.CommunityAsset, page int) ([]*model.CommunityAssetComment, error) {
	var comments []*model.CommunityAssetComment

	return comments, database.GetDatabase().Scopes(database.Paginate(page)).Where("community_asset_id = ?", obj.ID).Find(&comments).Error
}

func (r *communityAssetCommentResolver) User(ctx context.Context, obj *model.CommunityAssetComment) (*model.User, error) {
	user := new(model.User)

	return user, database.GetDatabase().First(user, obj.UserID).Error
}

func (r *discussionResolver) User(ctx context.Context, obj *model.Discussion) (*model.User, error) {
	user := new(model.User)

	return user, database.GetDatabase().First(user, obj.UserID).Error
}

func (r *discussionResolver) Game(ctx context.Context, obj *model.Discussion) (*model.Game, error) {
	game := new(model.Game)

	return game, database.GetDatabase().First(game, obj.GameID).Error
}

func (r *discussionResolver) Comments(ctx context.Context, obj *model.Discussion, page int) ([]*model.DiscussionComment, error) {
	var comments []*model.DiscussionComment

	return comments, database.GetDatabase().Scopes(database.Paginate(page)).Where("discussion_id = ?", obj.ID).Find(&comments).Error
}

func (r *discussionCommentResolver) User(ctx context.Context, obj *model.DiscussionComment) (*model.User, error) {
	user := new(model.User)

	return user, database.GetDatabase().First(user, obj.UserID).Error
}

func (r *friendRequestResolver) User(ctx context.Context, obj *model.FriendRequest) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *friendRequestResolver) Friend(ctx context.Context, obj *model.FriendRequest) (*model.Friends, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *friendsResolver) User(ctx context.Context, obj *model.Friends) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *friendsResolver) Friend(ctx context.Context, obj *model.Friends) (*model.Friends, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *gameResolver) GameBanner(ctx context.Context, obj *model.Game) (int, error) {
	return int(obj.GameGameBanner.ID), database.GetDatabase().Preload("GameGameBanner").First(obj).Error
}

func (r *gameResolver) GameSlideshow(ctx context.Context, obj *model.Game) ([]*model.GameMedia, error) {
	if err := database.GetDatabase().Preload("GameGameSlideshow").Preload("GameGameSlideshow.GameSlideshowMedia").First(obj).Error; err != nil {
		return nil, err
	}

	var gm []*model.GameMedia

	for _, es := range obj.GameGameSlideshow {
		gm = append(gm, &es.GameSlideshowMedia)
	}

	return gm, nil
}

func (r *gameResolver) Promo(ctx context.Context, obj *model.Game) (*model.Promo, error) {
	promo := model.Promo{}
	if err := database.GetDatabase().Debug().Where("game_id = ?", obj.ID).First(&promo).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	log.Print(promo)
	return &promo, nil
}

func (r *gameItemResolver) Game(ctx context.Context, obj *model.GameItem) (*model.Game, error) {
	game := new(model.Game)

	return game, database.GetDatabase().First(game, obj.GameID).Error
}

func (r *gameItemResolver) Transaction(ctx context.Context, obj *model.GameItem) ([]*model.MarketTransaction, error) {
	var tr []*model.MarketTransaction

	database.GetDatabase().Where("game_item_id = ? and created_at BETWEEN (now() - '1 month'::interval) and now()", obj.ID).Find(&tr)
	return tr, nil
}

func (r *gameMediaResolver) ContentType(ctx context.Context, obj *model.GameMedia) (string, error) {
	return obj.Type, nil
}

func (r *marketGameItemResolver) GameItem(ctx context.Context, obj *model.MarketGameItem) (*model.GameItem, error) {
	item := model.GameItem{}

	return &item, database.GetDatabase().First(&item, obj.GameItemID).Error
}

func (r *marketGameItemResolver) User(ctx context.Context, obj *model.MarketGameItem) (*model.User, error) {
	user := new(model.User)

	return user, database.GetDatabase().First(user, obj.UserID).Error
}

func (r *marketListingResolver) User(ctx context.Context, obj *model.MarketListing) (*model.User, error) {
	user := new(model.User)

	return user, database.GetDatabase().First(user, obj.UserID).Error
}

func (r *marketListingResolver) GameItem(ctx context.Context, obj *model.MarketListing) (*model.GameItem, error) {
	item := model.GameItem{}

	return &item, database.GetDatabase().First(&item, obj.GameItemID).Error
}

func (r *mutationResolver) Register(ctx context.Context, input model.Register) (*model.User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := model.User{
		AccountName: input.AccountName,
		Password:    string(hash),
		Points:      0,
		ProfilePic:  "default-profile-pic.jpg",
		DisplayName: input.AccountName,
		Wallet:      0,
	}

	return &user, database.GetDatabase().Create(&user).Error
}

func (r *mutationResolver) CreateGame(ctx context.Context, input model.NewGame) (*model.Game, error) {
	dataBanner, err := ioutil.ReadAll(input.GameBanner.File)
	if err != nil {
		return nil, err
	}

	banner := model.GameMedia{
		ImageVideo: dataBanner,
		Type:       input.GameBanner.ContentType,
	}

	game := model.Game{
		GameTitle:             input.GameTitle,
		GameDescription:       input.GameDescription,
		GamePrice:             input.GamePrice,
		GamePublisher:         input.GamePublisher,
		GameDeveloper:         input.GameDeveloper,
		GameTag:               input.GameTag,
		GameSystemRequirement: input.GameSystemRequirement,
		GameAdult:             input.GameAdult,
		GameGameBanner:        banner,
		MostHouredPlayed:      rand.Int63n(50),
	}

	return &game, database.GetDatabase().Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&game).Error; err != nil {
			return err
		}

		var slideshows []*model.GameSlideshow
		for _, s := range input.GameSlideshow {
			data, err := ioutil.ReadAll(s.File)
			if err != nil {
				return err
			}
			slideshow := model.GameSlideshow{
				GameSlideshowMedia: model.GameMedia{
					ImageVideo: data,
					Type:       s.ContentType,
				},
				GameGameID: game,
			}
			slideshows = append(slideshows, &slideshow)
		}

		return tx.Create(&slideshows).Error
	})
}

func (r *mutationResolver) DeleteGame(ctx context.Context, id int64) (*model.Game, error) {
	game := new(model.Game)
	if err := database.GetDatabase().First(game, id).Error; err != nil {
		return nil, err
	}
	if err := database.GetDatabase().
		Preload("GameGameSlideshow").
		Preload("GameGameBanner").
		Preload("GameGameSlideshow.GameSlideshowMedia").
		First(game).
		Error; err != nil {
		return nil, err
	}

	return game, database.GetDatabase().Transaction(func(tx *gorm.DB) error {
		for _, e := range game.GameGameSlideshow {
			if err := tx.Delete(&e).Error; err != nil {
				return err
			}

			if err := tx.Delete(&e.GameSlideshowMedia).Error; err != nil {
				return err
			}
		}

		if err := tx.Delete(game).Error; err != nil {
			return err
		}

		return tx.Delete(&game.GameGameBanner).Error
	})
}

func (r *mutationResolver) UpdateGame(ctx context.Context, input model.UpdateGame) (*model.Game, error) {
	game := new(model.Game)
	if err := database.GetDatabase().First(game, input.ID).Error; err != nil {
		return nil, err
	}

	if err := database.GetDatabase().
		Preload("GameGameSlideshow").
		Preload("GameGameBanner").
		Preload("GameGameSlideshow.GameSlideshowMedia").
		First(game).
		Error; err != nil {
		return nil, err
	}

	game.GameTitle = input.GameTitle
	game.GameDescription = input.GameDescription
	game.GamePrice = input.GamePrice
	game.GameTag = input.GameTag
	game.GameDeveloper = input.GameDeveloper
	game.GamePublisher = input.GamePublisher
	game.GameSystemRequirement = input.GameSystemRequirement
	game.GameAdult = input.GameAdult

	return game, database.GetDatabase().Transaction(func(tx *gorm.DB) error {
		if input.GameBanner != nil {
			bannerUpdate, err := ioutil.ReadAll(input.GameBanner.File)
			if err != nil {
				return err
			}

			game.GameGameBanner.ImageVideo = bannerUpdate
		}

		if input.GameSlideshow != nil {
			var slideshows []model.GameSlideshow

			for _, s := range input.GameSlideshow {
				slideshowUpdate, err := ioutil.ReadAll(s.File)
				if err != nil {
					return err
				}

				slideshows = append(slideshows, model.GameSlideshow{
					GameSlideshowMedia: model.GameMedia{
						ImageVideo: slideshowUpdate,
						Type:       s.ContentType,
					},
					GameGameID: *game,
				})
			}

			for _, s := range game.GameGameSlideshow {
				if err := tx.Delete(&s).Error; err != nil {
					return err
				}
			}

			game.GameGameSlideshow = slideshows
		}

		return tx.Save(game).Error
	})
}

func (r *mutationResolver) InsertPromo(ctx context.Context, input model.NewPromo) (*model.Promo, error) {
	game := model.Game{}

	if err := database.GetDatabase().First(&game, input.GameID).Error; err != nil {
		return nil, err
	}

	promo := model.Promo{
		Game_:         game,
		DiscountPromo: int64(input.DiscountPromo),
		EndDate:       input.EndDate,
	}

	return &promo, database.GetDatabase().Debug().Create(&promo).Error
}

func (r *mutationResolver) DeletePromo(ctx context.Context, id int64) (*model.Promo, error) {
	promo := model.Promo{}
	if err := database.GetDatabase().Where("game_id = ?", id).First(&promo).Error; err != nil {
		return nil, err
	}

	return &promo, database.GetDatabase().Delete(&promo).Error
}

func (r *mutationResolver) UpdatePromo(ctx context.Context, input model.NewPromo) (*model.Promo, error) {
	promo := model.Promo{}

	if err := database.GetDatabase().Where("game_id = ?", input.GameID).First(&promo).Error; err != nil {
		return nil, err
	}

	promo.DiscountPromo = int64(input.DiscountPromo)
	promo.EndDate = input.EndDate

	return &promo, database.GetDatabase().Save(&promo).Error
}

func (r *mutationResolver) InsertPointsItem(ctx context.Context, input model.NewPointItem) (*model.PointItem, error) {
	item := model.PointItem{
		ItemImg:    "",
		ItemPoints: 0,
		ItemType:   "",
	}

	return &item, database.GetDatabase().Debug().Create(&item).Error
}

func (r *mutationResolver) InsertCommunityAsset(ctx context.Context, input model.NewCommunityAsset) (*model.CommunityAsset, error) {
	asset := model.CommunityAsset{
		Asset:   "",
		Like:    0,
		Dislike: 0,
	}

	return &asset, database.GetDatabase().Create(&asset).Error
}

func (r *mutationResolver) LikeCommunityAsset(ctx context.Context, id int64) (*model.CommunityAsset, error) {
	asset := model.CommunityAsset{}

	if err := database.GetDatabase().First(&asset, id).Error; err != nil {
		return nil, err
	}

	asset.Like++
	return &asset, database.GetDatabase().Save(&asset).Error
}

func (r *mutationResolver) DislikeCommunityAsset(ctx context.Context, id int64) (*model.CommunityAsset, error) {
	asset := model.CommunityAsset{}

	if err := database.GetDatabase().First(&asset, id).Error; err != nil {
		return nil, err
	}

	asset.Dislike++
	return &asset, database.GetDatabase().Save(&asset).Error
}

func (r *mutationResolver) InsertCommunityComment(ctx context.Context, input *model.NewCommunityComment) (*model.CommunityAssetComment, error) {
	comment := model.CommunityAssetComment{
		UserID:           input.UserID,
		Comment:          input.Comment,
		CommunityAssetID: input.ID,
	}

	return &comment, database.GetDatabase().Debug().Create(&comment).Error
}

func (r *mutationResolver) InsertReview(ctx context.Context, input *model.NewReview) (*model.Review, error) {
	review := model.Review{
		UserID:      input.UserID,
		GameID:      input.GameID,
		Description: input.Description,
		Recommended: input.Recommended,
		Upvote:      0,
		Downvote:    0,
	}

	return &review, database.GetDatabase().Create(&review).Error
}

func (r *mutationResolver) InsertReviewComment(ctx context.Context, input *model.NewReviewComment) (*model.ReviewComment, error) {
	comment := model.ReviewComment{
		UserID:   input.UserID,
		Comment:  input.Comment,
		ReviewID: input.ID,
	}

	return &comment, database.GetDatabase().Create(&comment).Error
}

func (r *mutationResolver) InsertDiscussion(ctx context.Context, input *model.NewDiscussion) (*model.Discussion, error) {
	discussion := model.Discussion{
		UserID:      input.UserID,
		GameID:      input.GameID,
		Title:       input.Title,
		Description: input.Description,
	}

	return &discussion, database.GetDatabase().Create(&discussion).Error
}

func (r *mutationResolver) InsertDiscussionComment(ctx context.Context, input *model.NewDiscussionComment) (*model.DiscussionComment, error) {
	comment := model.DiscussionComment{
		UserID:       input.UserID,
		Comment:      input.Comment,
		DiscussionID: input.ID,
	}

	return &comment, database.GetDatabase().Create(&comment).Error
}

func (r *mutationResolver) InsertPointTransaction(ctx context.Context, userID int64, itemID int64) (bool, error) {
	transaction := model.PointShopTr{
		ItemID: itemID,
		UserID: userID,
	}

	item := model.Point_Item{}

	database.GetDatabase().First(&item, itemID)

	user := model.User{}

	database.GetDatabase().First(&user, userID)

	user.Points -= item.ItemPoints

	database.GetDatabase().Save(&user)

	return true, database.GetDatabase().Create(&transaction).Error
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input *model.UpdateUser) (*model.User, error) {
	user := new(model.User)
	if err := database.GetDatabase().First(user, input.ID).Error; err != nil {
		return nil, err
	}

	user.DisplayName = input.DisplayName
	user.RealName = input.RealName
	user.Country = input.Country
	user.CustomURL = input.CustomURL
	user.Summary = input.Summary

	return user, database.GetDatabase().Save(user).Error
}

func (r *mutationResolver) UpdateAvatar(ctx context.Context, id int64, profilePic string) (*model.User, error) {
	user := new(model.User)
	if err := database.GetDatabase().First(user, id).Error; err != nil {
		return nil, err
	}

	user.ProfilePic = profilePic

	return user, database.GetDatabase().Save(user).Error
}

func (r *mutationResolver) UpdateTheme(ctx context.Context, id int64, theme string) (*model.User, error) {
	user := new(model.User)
	if err := database.GetDatabase().First(user, id).Error; err != nil {
		return nil, err
	}

	user.Theme = theme

	return user, database.GetDatabase().Save(user).Error
}

func (r *mutationResolver) UpdateFrame(ctx context.Context, id int64, frameID int64) (*model.User, error) {
	user := new(model.User)
	if err := database.GetDatabase().First(user, id).Error; err != nil {
		return nil, err
	}

	user.FrameID = frameID

	return user, database.GetDatabase().Save(user).Error
}

func (r *mutationResolver) SendOtp(ctx context.Context, input int) (*int, error) {
	if err := generated.UseCache().Set(ctx, strconv.Itoa(input), input, 10*time.Second).Err(); err != nil {
		log.Fatal(err)
	}

	cached, _ := generated.UseCache().Get(ctx, strconv.Itoa(input)).Result()

	m := gomail.NewMessage()
	m.SetHeader("From", "gtheresandia@gmail.com")
	m.SetHeader("To", "gtheresandia@gmail.com")
	m.SetHeader("Subject", "OTP Code From Staem")
	m.SetBody("text", strconv.Itoa(input))

	d := gomail.NewDialer("smtp.gmail.com", 587, "gtheresandia@gmail.com", "gthzrlvgmsbbzenv")

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

	var cache, _ = strconv.Atoi(cached)

	return &cache, nil
}

func (r *mutationResolver) InsertRedeemCode(ctx context.Context, code string, amountMoney int) (*model.RedeemCode, error) {
	redeem := model.RedeemCode{
		Code:        "",
		MoneyAmount: 0,
	}

	return &redeem, database.GetDatabase().Create(&code).Error
}

func (r *mutationResolver) HelpfulReview(ctx context.Context, id int64) (*model.Review, error) {
	review := model.Review{}

	if err := database.GetDatabase().First(&review, id).Error; err != nil {
		return nil, err
	}

	review.Helpful++
	return &review, database.GetDatabase().Save(&review).Error
}

func (r *mutationResolver) NotHelpfulReview(ctx context.Context, id int64) (*model.Review, error) {
	review := model.Review{}

	if err := database.GetDatabase().First(&review, id).Error; err != nil {
		return nil, err
	}

	review.NotHelpful++
	return &review, database.GetDatabase().Save(&review).Error
}

func (r *mutationResolver) UpdateBackground(ctx context.Context, id int64, backgroundID int64) (*model.User, error) {
	user := new(model.User)
	if err := database.GetDatabase().First(user, id).Error; err != nil {
		return nil, err
	}

	user.BackgroundID = backgroundID

	return user, database.GetDatabase().Save(user).Error
}

func (r *mutationResolver) RedeemWalletCode(ctx context.Context, code string, userID int64) (*model.User, error) {
	redeem := model.RedeemCode{}

	if err := database.GetDatabase().Where("code = ?", code).First(&redeem).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	user := model.User{}

	database.GetDatabase().First(&user, userID)

	user.Wallet += redeem.MoneyAmount

	return &user, database.GetDatabase().Save(&user).Error
}

func (r *mutationResolver) UpdateBadge(ctx context.Context, id int64, badgeID int64) (*model.User, error) {
	user := new(model.User)
	if err := database.GetDatabase().First(user, id).Error; err != nil {
		return nil, err
	}

	user.BadgeID = badgeID

	return user, database.GetDatabase().Save(user).Error
}

func (r *mutationResolver) UpdateMiniBackground(ctx context.Context, id int64, miniBgID int64) (*model.User, error) {
	user := new(model.User)
	if err := database.GetDatabase().First(user, id).Error; err != nil {
		return nil, err
	}

	user.MiniBackgroundID = miniBgID

	return user, database.GetDatabase().Save(user).Error
}

func (r *mutationResolver) InsertBuyItem(ctx context.Context, userID int64, gameItemID int64, buyerID int64) (string, error) {
	marketGame := model.MarketGameItem{}

	temp := model.MarketGameItem{}

	database.GetDatabase().Where("user_id = ? and game_item_id = ? and type = ?", userID, gameItemID, "offer").First(&temp)

	database.GetDatabase().Create(&model.MarketTransaction{
		GameItemID: gameItemID,
		Price:      int(temp.Price),
	})

	//database.GetDatabase().Create(&model.Inventory{
	//	UserID:     userID,
	//	GameItemID: gameItemID,
	//})
	//
	//database.GetDatabase().Where("user_id = ? and game_item_id = ?", buyerID, gameItemID).Delete(&model.Inventory{})

	database.GetDatabase().Debug().Where("user_id = ? and game_item_id = ? and type = ? and price = ?", userID, gameItemID, "offer", temp.Price).Delete(&marketGame)

	m := gomail.NewMessage()
	m.SetHeader("From", "gtheresandia@gmail.com")
	m.SetHeader("To", "gtheresandia@gmail.com")
	m.SetHeader("Subject", "OTP Code From Staem")
	m.SetBody("text", "You successfully bought game item :D")

	d := gomail.NewDialer("smtp.gmail.com", 587, "gtheresandia@gmail.com", "gthzrlvgmsbbzenv")

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

	return "Success", nil
}

func (r *mutationResolver) InsertSellItem(ctx context.Context, userID int64, gameItemID int64, sellerID int64) (string, error) {
	marketGame := model.MarketGameItem{}

	temp := model.MarketGameItem{}

	database.GetDatabase().Where("user_id = ? and game_item_id = ? and type = ?", userID, gameItemID, "bid").First(&temp)

	database.GetDatabase().Create(&model.MarketTransaction{
		GameItemID: gameItemID,
		Price:      int(temp.Price),
	})

	//database.GetDatabase().Create(&model.Inventory{
	//	UserID:     sellerID,
	//	GameItemID: gameItemID,
	//})
	//
	//database.GetDatabase().Where("user_id = ? and game_item_id = ?", userID, gameItemID).Delete(&model.Inventory{})

	database.GetDatabase().Where("user_id = ? and game_item_id = ? and type = ? and price = ?", userID, gameItemID, "bid", temp.Price).Delete(&marketGame)

	m := gomail.NewMessage()
	m.SetHeader("From", "gtheresandia@gmail.com")
	m.SetHeader("To", "gtheresandia@gmail.com")
	m.SetHeader("Subject", "OTP Code From Staem")
	m.SetBody("text", "You successfully sell game item :D")

	d := gomail.NewDialer("smtp.gmail.com", 587, "gtheresandia@gmail.com", "gthzrlvgmsbbzenv")

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

	return "Success", nil
}

func (r *mutationResolver) InsertMarketItem(ctx context.Context, input model.NewMarketItem) (*model.MarketGameItem, error) {
	marketItem := model.MarketGameItem{
		GameItemID: input.GameItemID,
		UserID:     input.UserID,
		Price:      int64(input.Price),
		Type:       input.Type,
	}
	database.GetDatabase().Create(&model.MarketListing{
		UserID:     input.UserID,
		GameItemID: input.GameItemID,
		Price:      input.Price,
		Type:       input.Type,
	})

	var price = strconv.Itoa(input.Price)

	if input.Type == "offer" {
		for _, socket := range r.MarketSocket[int(input.GameItemID)] {
			socket <- "A user added offer for " + price
		}
	} else {
		for _, socket := range r.MarketSocket[int(input.GameItemID)] {
			socket <- "A user added bid for " + price
		}
	}

	return &marketItem, database.GetDatabase().Create(&marketItem).Error
}

func (r *mutationResolver) AddWalletAmount(ctx context.Context, userID int64, amount int) (*model.User, error) {
	user := &model.User{}

	database.GetDatabase().Where("id = ?", userID).Find(user)

	user.Wallet += int64(amount)

	return user, database.GetDatabase().Where("id = ?", userID).Save(user).Error
}

func (r *mutationResolver) ReduceWalletAmount(ctx context.Context, userID int64, amount int) (*model.User, error) {
	user := &model.User{}

	database.GetDatabase().Debug().Where("id = ?", userID).Find(user)

	user.Wallet -= int64(amount)

	return user, database.GetDatabase().Debug().Where("id = ?", userID).Save(user).Error
}

func (r *mutationResolver) InsertCommunityVidImg(ctx context.Context, imgVid string, userID int64) (*model.CommunityAsset, error) {
	asset := &model.CommunityAsset{
		Asset:   imgVid,
		Like:    0,
		Dislike: 0,
		UserID:  userID,
	}

	return asset, database.GetDatabase().Create(&asset).Error
}

func (r *mutationResolver) InsertNewReview(ctx context.Context, userID int64, gameID int64, desc string, recommend bool) (*model.Review, error) {
	review := &model.Review{
		UserID:      userID,
		GameID:      gameID,
		Description: desc,
		Recommended: recommend,
		Upvote:      0,
		Downvote:    0,
		Helpful:     0,
		NotHelpful:  0,
	}

	return review, database.GetDatabase().Create(&review).Error
}

func (r *mutationResolver) CreateUnsuspensionRequest(ctx context.Context, input model.InputUnsuspensionRequest) (*model.UnsuspensionRequest, error) {
	var user model.User
	database.GetDatabase().Where("email = ?", input.UserEmail).First(&user)

	result := model.UnsuspensionRequest{
		UserID: int(user.ID),
		Reason: input.Reason,
	}

	database.GetDatabase().Create(&result)

	return &result, nil
}

func (r *mutationResolver) CreateReportRequest(ctx context.Context, input model.InputRequestReport) (*model.ReportRequest, error) {
	var req = &model.ReportRequest{
		ReporterID:  input.ReporterID,
		SuspectedID: input.SuspectedID,
		Reason:      input.Reason,
	}

	database.GetDatabase().Create(&req)
	return req, nil
}

func (r *mutationResolver) AddReported(ctx context.Context, input int64) (*model.User, error) {
	var user model.User

	database.GetDatabase().Where("id = ?", input).First(&user)
	user.Reported++
	database.GetDatabase().Save(&user)

	return &user, nil
}

func (r *mutationResolver) CreateSuspensionList(ctx context.Context, input model.InputSuspensionList) (string, error) {
	var user model.User
	var ulist model.UnsuspensionRequest

	database.GetDatabase().Where("id = ?", input.UserID).First(&user)

	var list model.SuspensionList
	if input.Suspended {
		list = model.SuspensionList{
			UserID:    input.UserID,
			Reason:    input.Reason,
			Suspended: true,
		}
		user.Suspended = true
	} else {
		list = model.SuspensionList{
			UserID:    input.UserID,
			Reason:    input.Reason,
			Suspended: false,
		}
		user.Reported = 0
		user.Suspended = false
	}

	database.GetDatabase().Where("user_id = ?", input.UserID).Delete(&ulist)
	database.GetDatabase().Save(&user)
	database.GetDatabase().Create(&list)
	return "Success", nil
}

func (r *queryResolver) Login(ctx context.Context, accountName string, password string) (string, error) {
	user := model.User{}
	if err := database.GetDatabase().Where("account_name = ?", accountName).First(&user).Error; err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": user.ID,
	})

	return token.SignedString([]byte("skolastikagabriella"))
}

func (r *queryResolver) GameByID(ctx context.Context, id int64) (*model.Game, error) {
	game := new(model.Game)

	return game, database.GetDatabase().First(game, id).Error
}

func (r *queryResolver) GetGame(ctx context.Context) ([]*model.Game, error) {
	var games []*model.Game

	if err := database.GetDatabase().Find(&games).Error; err != nil {
		return nil, err
	}

	return games, nil
}

func (r *queryResolver) GetPromo(ctx context.Context, gameID int64) (*model.Promo, error) {
	promo := model.Promo{}

	return &promo, database.GetDatabase().Where("game_id = ?", gameID).First(&promo).Error
}

func (r *queryResolver) GetUser(ctx context.Context, jwtToken string) (*model.User, error) {
	user := model.User{}

	token, _ := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		return []byte("skolastikgabriella"), nil
	})

	database.GetDatabase().Where("id = ?", token.Claims.(jwt.MapClaims)["userID"]).Find(&user)

	return &user, nil
}

func (r *queryResolver) GetPointsItem(ctx context.Context) ([]*model.PointItem, error) {
	var points []*model.PointItem

	if err := database.GetDatabase().Find(&points).Error; err != nil {
		return nil, err
	}

	return points, nil
}

func (r *queryResolver) GetCommunityAsset(ctx context.Context) ([]*model.CommunityAsset, error) {
	var asset []*model.CommunityAsset

	if err := database.GetDatabase().Find(&asset).Error; err != nil {
		return nil, err
	}

	return asset, nil
}

func (r *queryResolver) GetCommunityAssetByID(ctx context.Context, id int64) (*model.CommunityAsset, error) {
	asset := new(model.CommunityAsset)

	return asset, database.GetDatabase().First(asset, id).Error
}

func (r *queryResolver) GetCommunityReview(ctx context.Context) ([]*model.Review, error) {
	var review []*model.Review

	if err := database.GetDatabase().Find(&review).Error; err != nil {
		return nil, err
	}

	return review, nil
}

func (r *queryResolver) GetReviewByID(ctx context.Context, id int64) (*model.Review, error) {
	review := new(model.Review)

	return review, database.GetDatabase().First(review, id).Error
}

func (r *queryResolver) GetDiscussion(ctx context.Context) ([]*model.Discussion, error) {
	var discussion []*model.Discussion

	if err := database.GetDatabase().Find(&discussion).Error; err != nil {
		return nil, err
	}

	return discussion, nil
}

func (r *queryResolver) GetDiscussionByID(ctx context.Context, id int64) (*model.Discussion, error) {
	discussion := new(model.Discussion)

	return discussion, database.GetDatabase().First(discussion, id).Error
}

func (r *queryResolver) GetRedeemCode(ctx context.Context, code string) (*model.RedeemCode, error) {
	redeem := model.RedeemCode{}

	return &redeem, database.GetDatabase().First(&redeem, code).Error
}

func (r *queryResolver) GetGameItem(ctx context.Context, page int) ([]*model.GameItem, error) {
	var items []*model.GameItem

	if err := database.GetDatabase().Scopes(database.Paginate(page)).Order("transaction_count desc").Find(&items).Error; err != nil {
		return nil, err
	}

	return items, nil
}

func (r *queryResolver) GetGameItemByID(ctx context.Context, id int64) (*model.GameItem, error) {
	gameItem := new(model.GameItem)

	return gameItem, database.GetDatabase().First(gameItem, id).Error
}

func (r *queryResolver) GetMarketGameItemByID(ctx context.Context, id int64) ([]*model.MarketGameItem, error) {
	var marketGameItem []*model.MarketGameItem

	if err := database.GetDatabase().Find(&marketGameItem, id).Error; err != nil {
		return nil, err
	}

	return marketGameItem, nil
}

func (r *queryResolver) GetMarketListing(ctx context.Context) ([]*model.MarketListing, error) {
	var marketGameListing []*model.MarketListing

	if err := database.GetDatabase().Find(&marketGameListing).Error; err != nil {
		return nil, err
	}

	return marketGameListing, nil
}

func (r *queryResolver) GetAllUser(ctx context.Context, page int) ([]*model.User, error) {
	var user []*model.User

	if err := database.GetDatabase().Scopes(database.Paginate(page)).Find(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *queryResolver) GetAllGame(ctx context.Context) ([]*model.Game, error) {
	var games []*model.Game

	if err := database.GetDatabase().Debug().Order("most_houred_played desc").Find(&games).Error; err != nil {
		return nil, err
	}

	return games, nil
}

func (r *queryResolver) GetReportRequest(ctx context.Context) ([]*model.ReportRequest, error) {
	var req []*model.ReportRequest

	database.GetDatabase().Find(&req)
	return req, nil
}

func (r *queryResolver) GetUnsuspensionRequest(ctx context.Context) ([]*model.UnsuspensionRequest, error) {
	var req []*model.UnsuspensionRequest

	database.GetDatabase().Find(&req)
	return req, nil
}

func (r *queryResolver) GetSuspensionList(ctx context.Context) ([]*model.SuspensionList, error) {
	var req []*model.SuspensionList

	database.GetDatabase().Find(&req)
	return req, nil
}

func (r *queryResolver) DeleteReport(ctx context.Context, id int64) (string, error) {
	var report model.ReportRequest

	database.GetDatabase().Where("id = ?", id).Delete(&report)
	return "Berhasil Delete", nil
}

func (r *reviewResolver) User(ctx context.Context, obj *model.Review) (*model.User, error) {
	user := model.User{}

	return &user, database.GetDatabase().First(&user, obj.UserID).Error
}

func (r *reviewResolver) Game(ctx context.Context, obj *model.Review) (*model.Game, error) {
	game := model.Game{}

	return &game, database.GetDatabase().First(&game, obj.GameID).Error
}

func (r *reviewResolver) Comments(ctx context.Context, obj *model.Review, page int) ([]*model.ReviewComment, error) {
	var comments []*model.ReviewComment

	return comments, database.GetDatabase().Scopes(database.Paginate(page)).Where("review_id = ?", obj.ID).Find(&comments).Error
}

func (r *reviewCommentResolver) User(ctx context.Context, obj *model.ReviewComment) (*model.User, error) {
	user := model.User{}

	return &user, database.GetDatabase().First(&user, obj.UserID).Error
}

func (r *subscriptionResolver) MessageAdded(ctx context.Context, itemID int) (<-chan string, error) {
	event := make(chan string, 1)
	r.MarketSocket[itemID] = append(r.MarketSocket[itemID], event)
	return event, nil
}

func (r *suspensionListResolver) User(ctx context.Context, obj *model.SuspensionList) (*model.User, error) {
	var user model.User
	database.GetDatabase().Where("id = ?", obj.UserID).First(&user)
	return &user, nil
}

func (r *unsuspensionRequestResolver) User(ctx context.Context, obj *model.UnsuspensionRequest) (*model.User, error) {
	var user model.User
	database.GetDatabase().Where("id = ?", obj.UserID).First(&user)
	return &user, nil
}

func (r *userResolver) Frame(ctx context.Context, obj *model.User) (*model.PointItem, error) {
	frame := model.PointItem{}

	database.GetDatabase().First(&frame, obj.FrameID)

	return &frame, nil
}

func (r *userResolver) OwnedFrame(ctx context.Context, obj *model.User) ([]*model.PointItem, error) {
	var frames []*model.PointItem

	return frames, database.GetDatabase().
		Joins("join point_shop_trs tr on point_items.id = tr.item_id").
		Where("user_id = ?", obj.ID).Where("item_type = 'Avatar Frame'").Find(&frames).Error
}

func (r *userResolver) Background(ctx context.Context, obj *model.User) (*model.PointItem, error) {
	background := model.PointItem{}

	database.GetDatabase().First(&background, obj.FrameID)

	return &background, nil
}

func (r *userResolver) OwnedBackground(ctx context.Context, obj *model.User) ([]*model.PointItem, error) {
	var backgrounds []*model.PointItem

	return backgrounds, database.GetDatabase().
		Joins("join point_shop_trs tr on point_items.id = tr.item_id").
		Where("user_id = ?", obj.ID).Where("item_type = 'Profile Background'").Find(&backgrounds).Error
}

func (r *userResolver) Badge(ctx context.Context, obj *model.User) (*model.PointItem, error) {
	badge := model.PointItem{}

	database.GetDatabase().First(&badge, obj.BadgeID)

	return &badge, nil
}

func (r *userResolver) OwnedBadge(ctx context.Context, obj *model.User) ([]*model.PointItem, error) {
	var badge []*model.PointItem

	return badge, database.GetDatabase().
		Joins("join point_shop_trs tr on point_items.id = tr.item_id").
		Where("user_id = ?", obj.ID).Where("item_type = 'Badge'").Find(&badge).Error
}

func (r *userResolver) MiniBackground(ctx context.Context, obj *model.User) (*model.PointItem, error) {
	miniBg := model.PointItem{}

	database.GetDatabase().First(&miniBg, obj.MiniBackgroundID)

	return &miniBg, nil
}

func (r *userResolver) OwnedMiniBackground(ctx context.Context, obj *model.User) ([]*model.PointItem, error) {
	var miniBg []*model.PointItem

	return miniBg, database.GetDatabase().Debug().
		Joins("join point_shop_trs tr on point_items.id = tr.item_id").
		Where("user_id = ?", obj.ID).Where("item_type = 'Mini Profile'").Find(&miniBg).Error
}

func (r *userResolver) Friends(ctx context.Context, obj *model.User) (*model.Friends, error) {
	panic("asdfasdfasdf")
}

// CommunityAsset returns generated.CommunityAssetResolver implementation.
func (r *Resolver) CommunityAsset() generated.CommunityAssetResolver {
	return &communityAssetResolver{r}
}

// CommunityAssetComment returns generated.CommunityAssetCommentResolver implementation.
func (r *Resolver) CommunityAssetComment() generated.CommunityAssetCommentResolver {
	return &communityAssetCommentResolver{r}
}

// Discussion returns generated.DiscussionResolver implementation.
func (r *Resolver) Discussion() generated.DiscussionResolver { return &discussionResolver{r} }

// DiscussionComment returns generated.DiscussionCommentResolver implementation.
func (r *Resolver) DiscussionComment() generated.DiscussionCommentResolver {
	return &discussionCommentResolver{r}
}

// FriendRequest returns generated.FriendRequestResolver implementation.
func (r *Resolver) FriendRequest() generated.FriendRequestResolver { return &friendRequestResolver{r} }

// Friends returns generated.FriendsResolver implementation.
func (r *Resolver) Friends() generated.FriendsResolver { return &friendsResolver{r} }

// Game returns generated.GameResolver implementation.
func (r *Resolver) Game() generated.GameResolver { return &gameResolver{r} }

// GameItem returns generated.GameItemResolver implementation.
func (r *Resolver) GameItem() generated.GameItemResolver { return &gameItemResolver{r} }

// GameMedia returns generated.GameMediaResolver implementation.
func (r *Resolver) GameMedia() generated.GameMediaResolver { return &gameMediaResolver{r} }

// MarketGameItem returns generated.MarketGameItemResolver implementation.
func (r *Resolver) MarketGameItem() generated.MarketGameItemResolver {
	return &marketGameItemResolver{r}
}

// MarketListing returns generated.MarketListingResolver implementation.
func (r *Resolver) MarketListing() generated.MarketListingResolver { return &marketListingResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Review returns generated.ReviewResolver implementation.
func (r *Resolver) Review() generated.ReviewResolver { return &reviewResolver{r} }

// ReviewComment returns generated.ReviewCommentResolver implementation.
func (r *Resolver) ReviewComment() generated.ReviewCommentResolver { return &reviewCommentResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

// SuspensionList returns generated.SuspensionListResolver implementation.
func (r *Resolver) SuspensionList() generated.SuspensionListResolver {
	return &suspensionListResolver{r}
}

// UnsuspensionRequest returns generated.UnsuspensionRequestResolver implementation.
func (r *Resolver) UnsuspensionRequest() generated.UnsuspensionRequestResolver {
	return &unsuspensionRequestResolver{r}
}

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type communityAssetResolver struct{ *Resolver }
type communityAssetCommentResolver struct{ *Resolver }
type discussionResolver struct{ *Resolver }
type discussionCommentResolver struct{ *Resolver }
type friendRequestResolver struct{ *Resolver }
type friendsResolver struct{ *Resolver }
type gameResolver struct{ *Resolver }
type gameItemResolver struct{ *Resolver }
type gameMediaResolver struct{ *Resolver }
type marketGameItemResolver struct{ *Resolver }
type marketListingResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type reviewResolver struct{ *Resolver }
type reviewCommentResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
type suspensionListResolver struct{ *Resolver }
type unsuspensionRequestResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
