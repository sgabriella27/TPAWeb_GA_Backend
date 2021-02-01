package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"io/ioutil"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/sgabriella27/TPAWebGA_Back/database"
	"github.com/sgabriella27/TPAWebGA_Back/graph/generated"
	"github.com/sgabriella27/TPAWebGA_Back/graph/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

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

func (r *gameMediaResolver) ContentType(ctx context.Context, obj *model.GameMedia) (string, error) {
	return obj.Type, nil
}

func (r *mutationResolver) Register(ctx context.Context, input model.Register) (*model.User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user := model.User{
		AccountName: input.AccountName,
		Password:    string(hash),
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

// Game returns generated.GameResolver implementation.
func (r *Resolver) Game() generated.GameResolver { return &gameResolver{r} }

// GameMedia returns generated.GameMediaResolver implementation.
func (r *Resolver) GameMedia() generated.GameMediaResolver { return &gameMediaResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type gameResolver struct{ *Resolver }
type gameMediaResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
