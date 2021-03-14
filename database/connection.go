package database

import (
	"github.com/sgabriella27/TPAWebGA_Back/graph/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"os"
)

var inst *gorm.DB

func init() {
	dsn := "host=localhost user=postgres password= dbname=tpaweb port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	if os.Getenv("DATABASE_URL") != "" {
		dsn = os.Getenv("DATABASE_URL")
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	inst = db

	migration()
	//seed()
	//seedPointItem()
	//seedCommunityAsset()
	//seedCommunityComment()
	//seedReview()
	//seedReviewComment()
	//seedDiscussion()
	//seedDiscussionComment()
	//seedPointTransaction()
	//seedProfileBg()
	//seedRedeemCode()
	//seedGameItem()
	//seedMarketGameItem()
}

func GetDatabase() *gorm.DB {
	return inst
}

func migration() {
	//if err := inst.Migrator().DropTable(&model.User{}, &model.Game{}, &model.GameMedia{}, &model.GameSlideshow{}); err != nil {
	//	log.Fatal(err)
	//}log.Fatal(err)

	if err := inst.AutoMigrate(&model.Promo{},
		&model.User{},
		&model.Game{},
		&model.GameMedia{},
		&model.GameSlideshow{},
		&model.Point_Item{},
		&model.CommunityAsset{},
		&model.CommunityAssetComment{},
		&model.Review{},
		&model.ReviewComment{},
		&model.Discussion{},
		&model.DiscussionComment{},
		&model.PointShopTr{},
		&model.RedeemCode{},
		&model.GameItem{},
		&model.MarketGameItem{},
		&model.MarketTransaction{},
		&model.MarketListing{},
		&model.Inventory{}); err != nil {
		log.Fatal(err)

	}
}

func seed() {
	hash, _ := bcrypt.GenerateFromPassword([]byte("gthere"), bcrypt.DefaultCost)
	//inst.Create(&model.User{
	//	AccountName: "admin",
	//	Password:    string(hash),
	//})
	inst.Create(&model.User{
		AccountName: "gthere",
		Password:    string(hash),
		Points:      999999,
		ProfilePic:  "panda.png",
		Wallet:      1000000000,
	})
}

func seedMarketGameItem() {
	inst.Create(&model.MarketGameItem{
		UserID:     41,
		GameItemID: 1,
		Price:      6000,
		Type:       "offer",
	})

	inst.Create(&model.MarketGameItem{
		UserID:     42,
		GameItemID: 1,
		Price:      8000,
		Type:       "offer",
	})

	inst.Create(&model.MarketGameItem{
		UserID:     44,
		GameItemID: 1,
		Price:      2700,
		Type:       "offer",
	})

	inst.Create(&model.MarketGameItem{
		UserID:     47,
		GameItemID: 1,
		Price:      3500,
		Type:       "bid",
	})

	inst.Create(&model.MarketGameItem{
		UserID:     48,
		GameItemID: 1,
		Price:      2100,
		Type:       "bid",
	})
}

func seedGameItem() {
	inst.Create(&model.GameItem{
		GameID:           2,
		GameItemName:     "Broken Fang",
		GameItemImg:      "https://community.cloudflare.steamstatic.com/economy/image/-9a81dlWLwJ2UUGcVs_nsVtzdOEdtWwKGZZLQHTxDZ7I56KU0Zwwo4NUX4oFJZEHLbXU5A1PIYQNqhpOSV-fRPasw8rsUFJ5KBFZv668FFU3naeZIWUStYjgxdnewfGmZb6DxW8AupMp27yT9IqiilCxqkRkZGyldoaLMlhp6IQjKcg/360fx360f",
		TransactionCount: rand.Int63n(50),
		GameItemDesc:     "ini game item desc nya seeding dipukul rata :D",
	})

	inst.Create(&model.GameItem{
		GameID:           2,
		GameItemName:     "Asimov",
		GameItemImg:      "https://community.cloudflare.steamstatic.com/economy/image/-9a81dlWLwJ2UUGcVs_nsVtzdOEdtWwKGZZLQHTxDZ7I56KU0Zwwo4NUX4oFJZEHLbXH5ApeO4YmlhxYQknCRvCo04DEVlxkKgpopL-zJAt21uH3cDx96t2ykb-GkuP1P7fYlVRD7dN-hv_E57P5gVO8vywwMiukcZjBdwBraVmG_1nsk-nug8fvus6YyHFj6HQm5HfdnUfliRFKbLE7habIVxzAUNH92sAX/360fx360f",
		TransactionCount: rand.Int63n(50),
		GameItemDesc:     "ini game item desc nya seeding dipukul rata :D",
	})

	inst.Create(&model.GameItem{
		GameID:           2,
		GameItemName:     "Digital Cased",
		GameItemImg:      "https://community.cloudflare.steamstatic.com/economy/image/-9a81dlWLwJ2UUGcVs_nsVtzdOEdtWwKGZZLQHTxDZ7I56KU0Zwwo4NUX4oFJZEHLbXH5ApeO4YmlhxYQknCRvCo04DEVlxkKgpou-6kejhjxszFJTwW09-vloWZh-L6OITck29Y_chOhujT8om72wy1-kBlYzryJI-UdAA8aAvU81e7w-zphJS06JrMnSdmvCkjtCrelgv33099jS-zpA/360fx360f",
		TransactionCount: rand.Int63n(50),
		GameItemDesc:     "ini game item desc nya seeding dipukul rata :D",
	})

	inst.Create(&model.GameItem{
		GameID:           2,
		GameItemName:     "Ready GO",
		GameItemImg:      "https://community.cloudflare.steamstatic.com/economy/image/-9a81dlWLwJ2UUGcVs_nsVtzdOEdtWwKGZZLQHTxDZ7I56KU0Zwwo4NUX4oFJZEHLbXH5ApeO4YmlhxYQknCRvCo04DEVlxkKgpot621FAR17PLfYQJD_9W7m5a0mvLwOq7cqWdQ-sJ0xOzAot-jiQa3-hBqYzvzLdSVJlQ3NQvR-FfsxL3qh5e7vM6bzSA26Sg8pSGKJUPeNtY/360fx360f",
		TransactionCount: rand.Int63n(50),
		GameItemDesc:     "ini game item desc nya seeding dipukul rata :D",
	})

	inst.Create(&model.GameItem{
		GameID:           2,
		GameItemName:     "Artificial Leak",
		GameItemImg:      "https://community.cloudflare.steamstatic.com/economy/image/-9a81dlWLwJ2UUGcVs_nsVtzdOEdtWwKGZZLQHTxDZ7I56KU0Zwwo4NUX4oFJZEHLbXQ9QVcJY8gulRfXkPbQuqS0c7dVBJ2Nwtcs7SaLQZu1MzAfjFN09C3hoeO2a72YO6HwzIH68Ek0--Uod3x2Qex-Es_NmmnJI7AcwBvZQzSrAS6k-zxxcjroZAOZ3Q/360fx360f",
		TransactionCount: rand.Int63n(50),
		GameItemDesc:     "ini game item desc nya seeding dipukul rata :D",
	})

	inst.Create(&model.GameItem{
		GameID:           2,
		GameItemName:     "Broken Fang",
		GameItemImg:      "https://community.cloudflare.steamstatic.com/economy/image/-9a81dlWLwJ2UUGcVs_nsVtzdOEdtWwKGZZLQHTxDZ7I56KU0Zwwo4NUX4oFJZEHLbXU5A1PIYQNqhpOSV-fRPasw8rsUFJ5KBFZv668FFU3naeZIWUStYjgxdnewfGmZb6DxW8AupMp27yT9IqiilCxqkRkZGyldoaLMlhp6IQjKcg/360fx360f",
		TransactionCount: rand.Int63n(50),
		GameItemDesc:     "ini game item desc nya seeding dipukul rata :D",
	})

	inst.Create(&model.GameItem{
		GameID:           2,
		GameItemName:     "Asimov",
		GameItemImg:      "https://community.cloudflare.steamstatic.com/economy/image/-9a81dlWLwJ2UUGcVs_nsVtzdOEdtWwKGZZLQHTxDZ7I56KU0Zwwo4NUX4oFJZEHLbXH5ApeO4YmlhxYQknCRvCo04DEVlxkKgpopL-zJAt21uH3cDx96t2ykb-GkuP1P7fYlVRD7dN-hv_E57P5gVO8vywwMiukcZjBdwBraVmG_1nsk-nug8fvus6YyHFj6HQm5HfdnUfliRFKbLE7habIVxzAUNH92sAX/360fx360f",
		TransactionCount: rand.Int63n(50),
		GameItemDesc:     "ini game item desc nya seeding dipukul rata :D",
	})

	inst.Create(&model.GameItem{
		GameID:           2,
		GameItemName:     "Digital Cased",
		GameItemImg:      "https://community.cloudflare.steamstatic.com/economy/image/-9a81dlWLwJ2UUGcVs_nsVtzdOEdtWwKGZZLQHTxDZ7I56KU0Zwwo4NUX4oFJZEHLbXH5ApeO4YmlhxYQknCRvCo04DEVlxkKgpou-6kejhjxszFJTwW09-vloWZh-L6OITck29Y_chOhujT8om72wy1-kBlYzryJI-UdAA8aAvU81e7w-zphJS06JrMnSdmvCkjtCrelgv33099jS-zpA/360fx360f",
		TransactionCount: rand.Int63n(50),
		GameItemDesc:     "ini game item desc nya seeding dipukul rata :D",
	})

	inst.Create(&model.GameItem{
		GameID:           2,
		GameItemName:     "Ready GO",
		GameItemImg:      "https://community.cloudflare.steamstatic.com/economy/image/-9a81dlWLwJ2UUGcVs_nsVtzdOEdtWwKGZZLQHTxDZ7I56KU0Zwwo4NUX4oFJZEHLbXH5ApeO4YmlhxYQknCRvCo04DEVlxkKgpot621FAR17PLfYQJD_9W7m5a0mvLwOq7cqWdQ-sJ0xOzAot-jiQa3-hBqYzvzLdSVJlQ3NQvR-FfsxL3qh5e7vM6bzSA26Sg8pSGKJUPeNtY/360fx360f",
		TransactionCount: rand.Int63n(50),
		GameItemDesc:     "ini game item desc nya seeding dipukul rata :D",
	})

	inst.Create(&model.GameItem{
		GameID:           2,
		GameItemName:     "Artificial Leak",
		GameItemImg:      "https://community.cloudflare.steamstatic.com/economy/image/-9a81dlWLwJ2UUGcVs_nsVtzdOEdtWwKGZZLQHTxDZ7I56KU0Zwwo4NUX4oFJZEHLbXQ9QVcJY8gulRfXkPbQuqS0c7dVBJ2Nwtcs7SaLQZu1MzAfjFN09C3hoeO2a72YO6HwzIH68Ek0--Uod3x2Qex-Es_NmmnJI7AcwBvZQzSrAS6k-zxxcjroZAOZ3Q/360fx360f",
		TransactionCount: rand.Int63n(50),
		GameItemDesc:     "ini game item desc nya seeding dipukul rata :D",
	})

	inst.Create(&model.GameItem{
		GameID:           3,
		GameItemName:     "Broken Fang",
		GameItemImg:      "https://community.cloudflare.steamstatic.com/economy/image/-9a81dlWLwJ2UUGcVs_nsVtzdOEdtWwKGZZLQHTxDZ7I56KU0Zwwo4NUX4oFJZEHLbXU5A1PIYQNqhpOSV-fRPasw8rsUFJ5KBFZv668FFU3naeZIWUStYjgxdnewfGmZb6DxW8AupMp27yT9IqiilCxqkRkZGyldoaLMlhp6IQjKcg/360fx360f",
		TransactionCount: rand.Int63n(50),
		GameItemDesc:     "ini game item desc nya seeding dipukul rata :D",
	})

	inst.Create(&model.GameItem{
		GameID:           3,
		GameItemName:     "Asimov",
		GameItemImg:      "https://community.cloudflare.steamstatic.com/economy/image/-9a81dlWLwJ2UUGcVs_nsVtzdOEdtWwKGZZLQHTxDZ7I56KU0Zwwo4NUX4oFJZEHLbXH5ApeO4YmlhxYQknCRvCo04DEVlxkKgpopL-zJAt21uH3cDx96t2ykb-GkuP1P7fYlVRD7dN-hv_E57P5gVO8vywwMiukcZjBdwBraVmG_1nsk-nug8fvus6YyHFj6HQm5HfdnUfliRFKbLE7habIVxzAUNH92sAX/360fx360f",
		TransactionCount: rand.Int63n(50),
		GameItemDesc:     "ini game item desc nya seeding dipukul rata :D",
	})

	inst.Create(&model.GameItem{
		GameID:           3,
		GameItemName:     "Digital Cased",
		GameItemImg:      "https://community.cloudflare.steamstatic.com/economy/image/-9a81dlWLwJ2UUGcVs_nsVtzdOEdtWwKGZZLQHTxDZ7I56KU0Zwwo4NUX4oFJZEHLbXH5ApeO4YmlhxYQknCRvCo04DEVlxkKgpou-6kejhjxszFJTwW09-vloWZh-L6OITck29Y_chOhujT8om72wy1-kBlYzryJI-UdAA8aAvU81e7w-zphJS06JrMnSdmvCkjtCrelgv33099jS-zpA/360fx360f",
		TransactionCount: rand.Int63n(50),
		GameItemDesc:     "ini game item desc nya seeding dipukul rata :D",
	})

	inst.Create(&model.GameItem{
		GameID:           3,
		GameItemName:     "Ready GO",
		GameItemImg:      "https://community.cloudflare.steamstatic.com/economy/image/-9a81dlWLwJ2UUGcVs_nsVtzdOEdtWwKGZZLQHTxDZ7I56KU0Zwwo4NUX4oFJZEHLbXH5ApeO4YmlhxYQknCRvCo04DEVlxkKgpot621FAR17PLfYQJD_9W7m5a0mvLwOq7cqWdQ-sJ0xOzAot-jiQa3-hBqYzvzLdSVJlQ3NQvR-FfsxL3qh5e7vM6bzSA26Sg8pSGKJUPeNtY/360fx360f",
		TransactionCount: rand.Int63n(50),
		GameItemDesc:     "ini game item desc nya seeding dipukul rata :D",
	})

	inst.Create(&model.GameItem{
		GameID:           3,
		GameItemName:     "Artificial Leak",
		GameItemImg:      "https://community.cloudflare.steamstatic.com/economy/image/-9a81dlWLwJ2UUGcVs_nsVtzdOEdtWwKGZZLQHTxDZ7I56KU0Zwwo4NUX4oFJZEHLbXQ9QVcJY8gulRfXkPbQuqS0c7dVBJ2Nwtcs7SaLQZu1MzAfjFN09C3hoeO2a72YO6HwzIH68Ek0--Uod3x2Qex-Es_NmmnJI7AcwBvZQzSrAS6k-zxxcjroZAOZ3Q/360fx360f",
		TransactionCount: rand.Int63n(50),
		GameItemDesc:     "ini game item desc nya seeding dipukul rata :D",
	})

	inst.Create(&model.GameItem{
		GameID:           4,
		GameItemName:     "Broken Fang",
		GameItemImg:      "https://community.cloudflare.steamstatic.com/economy/image/-9a81dlWLwJ2UUGcVs_nsVtzdOEdtWwKGZZLQHTxDZ7I56KU0Zwwo4NUX4oFJZEHLbXU5A1PIYQNqhpOSV-fRPasw8rsUFJ5KBFZv668FFU3naeZIWUStYjgxdnewfGmZb6DxW8AupMp27yT9IqiilCxqkRkZGyldoaLMlhp6IQjKcg/360fx360f",
		TransactionCount: rand.Int63n(50),
		GameItemDesc:     "ini game item desc nya seeding dipukul rata :D",
	})

	inst.Create(&model.GameItem{
		GameID:           4,
		GameItemName:     "Asimov",
		GameItemImg:      "https://community.cloudflare.steamstatic.com/economy/image/-9a81dlWLwJ2UUGcVs_nsVtzdOEdtWwKGZZLQHTxDZ7I56KU0Zwwo4NUX4oFJZEHLbXH5ApeO4YmlhxYQknCRvCo04DEVlxkKgpopL-zJAt21uH3cDx96t2ykb-GkuP1P7fYlVRD7dN-hv_E57P5gVO8vywwMiukcZjBdwBraVmG_1nsk-nug8fvus6YyHFj6HQm5HfdnUfliRFKbLE7habIVxzAUNH92sAX/360fx360f",
		TransactionCount: rand.Int63n(50),
		GameItemDesc:     "ini game item desc nya seeding dipukul rata :D",
	})

	inst.Create(&model.GameItem{
		GameID:           4,
		GameItemName:     "Digital Cased",
		GameItemImg:      "https://community.cloudflare.steamstatic.com/economy/image/-9a81dlWLwJ2UUGcVs_nsVtzdOEdtWwKGZZLQHTxDZ7I56KU0Zwwo4NUX4oFJZEHLbXH5ApeO4YmlhxYQknCRvCo04DEVlxkKgpou-6kejhjxszFJTwW09-vloWZh-L6OITck29Y_chOhujT8om72wy1-kBlYzryJI-UdAA8aAvU81e7w-zphJS06JrMnSdmvCkjtCrelgv33099jS-zpA/360fx360f",
		TransactionCount: rand.Int63n(50),
		GameItemDesc:     "ini game item desc nya seeding dipukul rata :D",
	})

	inst.Create(&model.GameItem{
		GameID:           4,
		GameItemName:     "Ready GO",
		GameItemImg:      "https://community.cloudflare.steamstatic.com/economy/image/-9a81dlWLwJ2UUGcVs_nsVtzdOEdtWwKGZZLQHTxDZ7I56KU0Zwwo4NUX4oFJZEHLbXH5ApeO4YmlhxYQknCRvCo04DEVlxkKgpot621FAR17PLfYQJD_9W7m5a0mvLwOq7cqWdQ-sJ0xOzAot-jiQa3-hBqYzvzLdSVJlQ3NQvR-FfsxL3qh5e7vM6bzSA26Sg8pSGKJUPeNtY/360fx360f",
		TransactionCount: rand.Int63n(50),
		GameItemDesc:     "ini game item desc nya seeding dipukul rata :D",
	})

	inst.Create(&model.GameItem{
		GameID:           4,
		GameItemName:     "Artificial Leak",
		GameItemImg:      "https://community.cloudflare.steamstatic.com/economy/image/-9a81dlWLwJ2UUGcVs_nsVtzdOEdtWwKGZZLQHTxDZ7I56KU0Zwwo4NUX4oFJZEHLbXQ9QVcJY8gulRfXkPbQuqS0c7dVBJ2Nwtcs7SaLQZu1MzAfjFN09C3hoeO2a72YO6HwzIH68Ek0--Uod3x2Qex-Es_NmmnJI7AcwBvZQzSrAS6k-zxxcjroZAOZ3Q/360fx360f",
		TransactionCount: rand.Int63n(50),
		GameItemDesc:     "ini game item desc nya seeding dipukul rata :D",
	})

	inst.Create(&model.GameItem{
		GameID:           4,
		GameItemName:     "Broken Fang",
		GameItemImg:      "https://community.cloudflare.steamstatic.com/economy/image/-9a81dlWLwJ2UUGcVs_nsVtzdOEdtWwKGZZLQHTxDZ7I56KU0Zwwo4NUX4oFJZEHLbXU5A1PIYQNqhpOSV-fRPasw8rsUFJ5KBFZv668FFU3naeZIWUStYjgxdnewfGmZb6DxW8AupMp27yT9IqiilCxqkRkZGyldoaLMlhp6IQjKcg/360fx360f",
		TransactionCount: rand.Int63n(50),
		GameItemDesc:     "ini game item desc nya seeding dipukul rata :D",
	})

	inst.Create(&model.GameItem{
		GameID:           4,
		GameItemName:     "Asimov",
		GameItemImg:      "https://community.cloudflare.steamstatic.com/economy/image/-9a81dlWLwJ2UUGcVs_nsVtzdOEdtWwKGZZLQHTxDZ7I56KU0Zwwo4NUX4oFJZEHLbXH5ApeO4YmlhxYQknCRvCo04DEVlxkKgpopL-zJAt21uH3cDx96t2ykb-GkuP1P7fYlVRD7dN-hv_E57P5gVO8vywwMiukcZjBdwBraVmG_1nsk-nug8fvus6YyHFj6HQm5HfdnUfliRFKbLE7habIVxzAUNH92sAX/360fx360f",
		TransactionCount: rand.Int63n(50),
		GameItemDesc:     "ini game item desc nya seeding dipukul rata :D",
	})

	inst.Create(&model.GameItem{
		GameID:           4,
		GameItemName:     "Digital Cased",
		GameItemImg:      "https://community.cloudflare.steamstatic.com/economy/image/-9a81dlWLwJ2UUGcVs_nsVtzdOEdtWwKGZZLQHTxDZ7I56KU0Zwwo4NUX4oFJZEHLbXH5ApeO4YmlhxYQknCRvCo04DEVlxkKgpou-6kejhjxszFJTwW09-vloWZh-L6OITck29Y_chOhujT8om72wy1-kBlYzryJI-UdAA8aAvU81e7w-zphJS06JrMnSdmvCkjtCrelgv33099jS-zpA/360fx360f",
		TransactionCount: rand.Int63n(50),
		GameItemDesc:     "ini game item desc nya seeding dipukul rata :D",
	})

	inst.Create(&model.GameItem{
		GameID:           4,
		GameItemName:     "Ready GO",
		GameItemImg:      "https://community.cloudflare.steamstatic.com/economy/image/-9a81dlWLwJ2UUGcVs_nsVtzdOEdtWwKGZZLQHTxDZ7I56KU0Zwwo4NUX4oFJZEHLbXH5ApeO4YmlhxYQknCRvCo04DEVlxkKgpot621FAR17PLfYQJD_9W7m5a0mvLwOq7cqWdQ-sJ0xOzAot-jiQa3-hBqYzvzLdSVJlQ3NQvR-FfsxL3qh5e7vM6bzSA26Sg8pSGKJUPeNtY/360fx360f",
		TransactionCount: rand.Int63n(50),
		GameItemDesc:     "ini game item desc nya seeding dipukul rata :D",
	})

	inst.Create(&model.GameItem{
		GameID:           4,
		GameItemName:     "Artificial Leak",
		GameItemImg:      "https://community.cloudflare.steamstatic.com/economy/image/-9a81dlWLwJ2UUGcVs_nsVtzdOEdtWwKGZZLQHTxDZ7I56KU0Zwwo4NUX4oFJZEHLbXQ9QVcJY8gulRfXkPbQuqS0c7dVBJ2Nwtcs7SaLQZu1MzAfjFN09C3hoeO2a72YO6HwzIH68Ek0--Uod3x2Qex-Es_NmmnJI7AcwBvZQzSrAS6k-zxxcjroZAOZ3Q/360fx360f",
		TransactionCount: rand.Int63n(50),
		GameItemDesc:     "ini game item desc nya seeding dipukul rata :D",
	})
}

func seedRedeemCode() {
	inst.Create(&model.RedeemCode{
		Code:        "abcd1",
		MoneyAmount: 50000,
	})
	inst.Create(&model.RedeemCode{
		Code:        "abcd2",
		MoneyAmount: 100000,
	})
	inst.Create(&model.RedeemCode{
		Code:        "abcd3",
		MoneyAmount: 150000,
	})
	inst.Create(&model.RedeemCode{
		Code:        "abcd4",
		MoneyAmount: 200000,
	})
	inst.Create(&model.RedeemCode{
		Code:        "abcd5",
		MoneyAmount: 300000,
	})
}

func seedPointTransaction() {
	inst.Create(&model.PointShopTr{
		ItemID: 139,
		UserID: 44,
	})
}

func seedReviewComment() {
	inst.Create(&model.ReviewComment{
		UserID:   2,
		Comment:  "k",
		ReviewID: 1,
	})
}

func seedDiscussionComment() {
	inst.Create(&model.DiscussionComment{
		UserID:       41,
		Comment:      "hallo semuanya",
		DiscussionID: 4,
	})
	inst.Create(&model.DiscussionComment{
		UserID:       42,
		Comment:      "hallo semuanya1",
		DiscussionID: 4,
	})
	inst.Create(&model.DiscussionComment{
		UserID:       42,
		Comment:      "hallo semuanya2",
		DiscussionID: 4,
	})
	inst.Create(&model.DiscussionComment{
		UserID:       41,
		Comment:      "hallo semuanya3",
		DiscussionID: 4,
	})
}

func seedDiscussion() {
	inst.Create(&model.Discussion{
		UserID:      41,
		GameID:      3,
		Title:       "Discussion Baru",
		Description: "lalalalalalalal",
	})
	inst.Create(&model.Discussion{
		UserID:      2,
		GameID:      3,
		Title:       "Discussion Baru2",
		Description: "lalalalalalalal2a",
	})
	inst.Create(&model.Discussion{
		UserID:      42,
		GameID:      3,
		Title:       "Discussion Baru3",
		Description: "lalalalalalalal3",
	})
}

func seedReview() {
	//inst.Create(&model.Review{
	//	UserID:      1,
	//	GameID:      2,
	//	Description: "Game ini sangat bagus",
	//	Recommended: true,
	//	Upvote:      10,
	//	Downvote:    5,
	//})
	//inst.Create(&model.Review{
	//	UserID:      2,
	//	GameID:      3,
	//	Description: "Game ini sangat seram",
	//	Recommended: false,
	//	Upvote:      20,
	//	Downvote:    4,
	//})
	//inst.Create(&model.Review{
	//	UserID:      2,
	//	GameID:      4,
	//	Description: "Game ini sangat seru",
	//	Recommended: true,
	//	Upvote:      50,
	//	Downvote:    1,
	//})
	inst.Create(&model.Review{
		UserID:      2,
		GameID:      4,
		Description: "Game ini sangat seru",
		Recommended: true,
		Upvote:      50,
		Downvote:    1,
		Helpful:     2,
		NotHelpful:  1,
	})
}

func seedCommunityAsset() {
	inst.Create(&model.CommunityAsset{
		Asset:   "image2.jpg",
		Like:    20,
		Dislike: 100,
		UserID:  1,
	})
}

func seedCommunityComment() {
	inst.Create(&model.CommunityAssetComment{
		UserID:           2,
		Comment:          "aku suka kamu love",
		CommunityAssetID: 1,
	})
	inst.Create(&model.CommunityAssetComment{
		UserID:           2,
		Comment:          "aku suka dia love",
		CommunityAssetID: 1,
	})
	inst.Create(&model.CommunityAssetComment{
		UserID:           2,
		Comment:          "aku suka kamu bgt",
		CommunityAssetID: 1,
	})
	inst.Create(&model.CommunityAssetComment{
		UserID:           2,
		Comment:          "kamu suka dia",
		CommunityAssetID: 1,
	})
	inst.Create(&model.CommunityAssetComment{
		UserID:           2,
		Comment:          "dia suka aku love",
		CommunityAssetID: 1,
	})
	inst.Create(&model.CommunityAssetComment{
		UserID:           2,
		Comment:          "aku suka dia",
		CommunityAssetID: 1,
	})
	inst.Create(&model.CommunityAssetComment{
		UserID:           2,
		Comment:          "aku gasuka kamu",
		CommunityAssetID: 1,
	})
	inst.Create(&model.CommunityAssetComment{
		UserID:           2,
		Comment:          "kamu gasuka aku",
		CommunityAssetID: 1,
	})
	inst.Create(&model.CommunityAssetComment{
		UserID:           2,
		Comment:          "dia gasuka kamu",
		CommunityAssetID: 1,
	})
	inst.Create(&model.CommunityAssetComment{
		UserID:           2,
		Comment:          "bodo amat",
		CommunityAssetID: 1,
	})
	inst.Create(&model.CommunityAssetComment{
		UserID:           2,
		Comment:          "ini comment dummy",
		CommunityAssetID: 1,
	})
}

func seedProfileBg() {
	inst.Create(&model.Point_Item{
		ItemImg:    "profile-background1.jpg",
		ItemPoints: 2000,
		ItemType:   "Profile Background",
	})
	inst.Create(&model.Point_Item{
		ItemImg:    "profile-background2.jpg",
		ItemPoints: 1500,
		ItemType:   "Profile Background",
	})
	inst.Create(&model.Point_Item{
		ItemImg:    "profile-background3.jpg",
		ItemPoints: 3000,
		ItemType:   "Profile Background",
	})
	inst.Create(&model.Point_Item{
		ItemImg:    "profile-background4.jpg",
		ItemPoints: 2100,
		ItemType:   "Profile Background",
	})
	inst.Create(&model.Point_Item{
		ItemImg:    "profile-background5.jpg",
		ItemPoints: 3500,
		ItemType:   "Profile Background",
	})
}

func seedPointItem() {
	inst.Create(&model.Point_Item{
		ItemImg:    "frame1.png",
		ItemPoints: 1300,
		ItemType:   "Avatar Frame",
	})
	inst.Create(&model.Point_Item{
		ItemImg:    "frame2.png",
		ItemPoints: 2500,
		ItemType:   "Avatar Frame",
	})
	inst.Create(&model.Point_Item{
		ItemImg:    "frame3.png",
		ItemPoints: 4000,
		ItemType:   "Avatar Frame",
	})
	inst.Create(&model.Point_Item{
		ItemImg:    "frame4.png",
		ItemPoints: 4600,
		ItemType:   "Avatar Frame",
	})
	inst.Create(&model.Point_Item{
		ItemImg:    "frame5.png",
		ItemPoints: 2500,
		ItemType:   "Avatar Frame",
	})
	inst.Create(&model.Point_Item{
		ItemImg:    "frame6.png",
		ItemPoints: 1700,
		ItemType:   "Avatar Frame",
	})
	inst.Create(&model.Point_Item{
		ItemImg:    "frame7.png",
		ItemPoints: 3100,
		ItemType:   "Avatar Frame",
	})
	inst.Create(&model.Point_Item{
		ItemImg:    "frame8.png",
		ItemPoints: 2400,
		ItemType:   "Avatar Frame",
	})
	inst.Create(&model.Point_Item{
		ItemImg:    "frame9.png",
		ItemPoints: 2800,
		ItemType:   "Avatar Frame",
	})
	inst.Create(&model.Point_Item{
		ItemImg:    "frame10.png",
		ItemPoints: 1600,
		ItemType:   "Avatar Frame",
	})
	inst.Create(&model.Point_Item{
		ItemImg:    "frame11.png",
		ItemPoints: 3200,
		ItemType:   "Avatar Frame",
	})
	inst.Create(&model.Point_Item{
		ItemImg:    "frame12.png",
		ItemPoints: 2200,
		ItemType:   "Avatar Frame",
	})
	inst.Create(&model.Point_Item{
		ItemImg:    "badge1.png",
		ItemPoints: 1000,
		ItemType:   "Badge",
	})
	inst.Create(&model.Point_Item{
		ItemImg:    "badge2.png",
		ItemPoints: 2200,
		ItemType:   "Badge",
	})
	inst.Create(&model.Point_Item{
		ItemImg:    "badge3.png",
		ItemPoints: 3000,
		ItemType:   "Badge",
	})
	inst.Create(&model.Point_Item{
		ItemImg:    "badge4.png",
		ItemPoints: 2100,
		ItemType:   "Badge",
	})
	inst.Create(&model.Point_Item{
		ItemImg:    "badge5.png",
		ItemPoints: 4200,
		ItemType:   "Badge",
	})
	inst.Create(&model.Point_Item{
		ItemImg:    "badge6.png",
		ItemPoints: 1500,
		ItemType:   "Badge",
	})
	inst.Create(&model.Point_Item{
		ItemImg:    "badge7.png",
		ItemPoints: 3100,
		ItemType:   "Badge",
	})
	inst.Create(&model.Point_Item{
		ItemImg:    "badge8.png",
		ItemPoints: 2400,
		ItemType:   "Badge",
	})
	inst.Create(&model.Point_Item{
		ItemImg:    "badge9.png",
		ItemPoints: 2700,
		ItemType:   "Badge",
	})
	inst.Create(&model.Point_Item{
		ItemImg:    "badge10.png",
		ItemPoints: 1600,
		ItemType:   "Badge",
	})
	inst.Create(&model.Point_Item{
		ItemImg:    "sticker1.png",
		ItemPoints: 1000,
		ItemType:   "Chat Sticker",
	})
	inst.Create(&model.Point_Item{
		ItemImg:    "sticker2.png",
		ItemPoints: 1900,
		ItemType:   "Chat Sticker",
	})
	inst.Create(&model.Point_Item{
		ItemImg:    "sticker3.png",
		ItemPoints: 3000,
		ItemType:   "Chat Sticker",
	})
	inst.Create(&model.Point_Item{
		ItemImg:    "sticker4.png",
		ItemPoints: 2100,
		ItemType:   "Chat Sticker",
	})
	inst.Create(&model.Point_Item{
		ItemImg:    "sticker5.png",
		ItemPoints: 1200,
		ItemType:   "Chat Sticker",
	})
	inst.Create(&model.Point_Item{
		ItemImg:    "sticker6.png",
		ItemPoints: 1500,
		ItemType:   "Chat Sticker",
	})
	inst.Create(&model.Point_Item{
		ItemImg:    "sticker7.png",
		ItemPoints: 3100,
		ItemType:   "Chat Sticker",
	})
	inst.Create(&model.Point_Item{
		ItemImg:    "sticker8.png",
		ItemPoints: 2400,
		ItemType:   "Chat Sticker",
	})
	inst.Create(&model.Point_Item{
		ItemImg:    "sticker9.png",
		ItemPoints: 2700,
		ItemType:   "Chat Sticker",
	})
	inst.Create(&model.Point_Item{
		ItemImg:    "sticker10.png",
		ItemPoints: 4000,
		ItemType:   "Chat Sticker",
	})
	inst.Create(&model.Point_Item{
		ItemImg:    "sticker11.png",
		ItemPoints: 3200,
		ItemType:   "Chat Sticker",
	})
	inst.Create(&model.Point_Item{
		ItemImg:    "sticker12.png",
		ItemPoints: 2200,
		ItemType:   "Chat Sticker",
	})
	inst.Create(&model.Point_Item{
		ItemImg:    "item1.png",
		ItemPoints: 1000,
		ItemType:   "Game Item",
	})
	inst.Create(&model.Point_Item{
		ItemImg:    "item2.png",
		ItemPoints: 2200,
		ItemType:   "Game Item",
	})
	inst.Create(&model.Point_Item{
		ItemImg:    "item3.png",
		ItemPoints: 3000,
		ItemType:   "Game Item",
	})
	inst.Create(&model.Point_Item{
		ItemImg:    "item4.png",
		ItemPoints: 2100,
		ItemType:   "Game Item",
	})
	inst.Create(&model.Point_Item{
		ItemImg:    "item5.png",
		ItemPoints: 1200,
		ItemType:   "Game Item",
	})
	inst.Create(&model.Point_Item{
		ItemImg:    "item6.png",
		ItemPoints: 1500,
		ItemType:   "Game Item",
	})
	inst.Create(&model.Point_Item{
		ItemImg:    "item7.png",
		ItemPoints: 3100,
		ItemType:   "Game Item",
	})
	inst.Create(&model.Point_Item{
		ItemImg:    "item8.png",
		ItemPoints: 2400,
		ItemType:   "Game Item",
	})
	inst.Create(&model.Point_Item{
		ItemImg:    "item9.png",
		ItemPoints: 2700,
		ItemType:   "Game Item",
	})
	inst.Create(&model.Point_Item{
		ItemImg:    "item10.png",
		ItemPoints: 1600,
		ItemType:   "Game Item",
	})
	inst.Create(&model.Point_Item{
		ItemImg:    "item11.png",
		ItemPoints: 3200,
		ItemType:   "Game Item",
	})
	inst.Create(&model.Point_Item{
		ItemImg:    "item12.png",
		ItemPoints: 2200,
		ItemType:   "Game Item",
	})
	inst.Create(&model.Point_Item{
		ItemImg:    "mini-profile1.jpg",
		ItemPoints: 1000,
		ItemType:   "Mini Profile",
	})
	inst.Create(&model.Point_Item{
		ItemImg:    "mini-profile2.jpg",
		ItemPoints: 2200,
		ItemType:   "Mini Profile",
	})
	inst.Create(&model.Point_Item{
		ItemImg:    "mini-profile3.jpg",
		ItemPoints: 3000,
		ItemType:   "Mini Profile",
	})
	inst.Create(&model.Point_Item{
		ItemImg:    "mini-profile4.jpg",
		ItemPoints: 2100,
		ItemType:   "Mini Profile",
	})
	inst.Create(&model.Point_Item{
		ItemImg:    "mini-profile5.jpg",
		ItemPoints: 1200,
		ItemType:   "Mini Profile",
	})
	inst.Create(&model.Point_Item{
		ItemImg:    "mini-profile6.jpg",
		ItemPoints: 1500,
		ItemType:   "Mini Profile",
	})
	inst.Create(&model.Point_Item{
		ItemImg:    "mini-profile7.jpg",
		ItemPoints: 3100,
		ItemType:   "Mini Profile",
	})
	inst.Create(&model.Point_Item{
		ItemImg:    "mini-profile8.jpg",
		ItemPoints: 2400,
		ItemType:   "Mini Profile",
	})
	inst.Create(&model.Point_Item{
		ItemImg:    "mini-profile9.jpg",
		ItemPoints: 2700,
		ItemType:   "Mini Profile",
	})
	inst.Create(&model.Point_Item{
		ItemImg:    "mini-profile10.jpg",
		ItemPoints: 1600,
		ItemType:   "Mini Profile",
	})
	inst.Create(&model.Point_Item{
		ItemImg:    "mini-profile11.jpg",
		ItemPoints: 3200,
		ItemType:   "Mini Profile",
	})
	inst.Create(&model.Point_Item{
		ItemImg:    "mini-profile12.jpg",
		ItemPoints: 2200,
		ItemType:   "Mini Profile",
	})
	inst.Create(&model.Point_Item{
		ItemImg:    "avatar-animated1.gif",
		ItemPoints: 1000,
		ItemType:   "Animated Avatar",
	})
	inst.Create(&model.Point_Item{
		ItemImg:    "avatar-animated2.gif",
		ItemPoints: 2200,
		ItemType:   "Animated Avatar",
	})
	inst.Create(&model.Point_Item{
		ItemImg:    "avatar-animated3.gif",
		ItemPoints: 5000,
		ItemType:   "Animated Avatar",
	})
	inst.Create(&model.Point_Item{
		ItemImg:    "avatar-animated4.gif",
		ItemPoints: 2100,
		ItemType:   "Animated Avatar",
	})
	inst.Create(&model.Point_Item{
		ItemImg:    "avatar-animated5.gif",
		ItemPoints: 2300,
		ItemType:   "Animated Avatar",
	})
	inst.Create(&model.Point_Item{
		ItemImg:    "avatar-animated6.gif",
		ItemPoints: 1500,
		ItemType:   "Animated Avatar",
	})
	inst.Create(&model.Point_Item{
		ItemImg:    "avatar-animated7.gif",
		ItemPoints: 2400,
		ItemType:   "Animated Avatar",
	})
	inst.Create(&model.Point_Item{
		ItemImg:    "avatar-animated8.gif",
		ItemPoints: 2400,
		ItemType:   "Animated Avatar",
	})
	inst.Create(&model.Point_Item{
		ItemImg:    "avatar-animated9.gif",
		ItemPoints: 2700,
		ItemType:   "Animated Avatar",
	})
	inst.Create(&model.Point_Item{
		ItemImg:    "avatar-animated10.gif",
		ItemPoints: 2000,
		ItemType:   "Animated Avatar",
	})
	inst.Create(&model.Point_Item{
		ItemImg:    "avatar-animated11.gif",
		ItemPoints: 5000,
		ItemType:   "Animated Avatar",
	})
	inst.Create(&model.Point_Item{
		ItemImg:    "avatar-animated12.gif",
		ItemPoints: 1900,
		ItemType:   "Animated Avatar",
	})
}
