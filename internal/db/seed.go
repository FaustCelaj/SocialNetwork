package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"strings"

	"github.com/faustcelaj/social_project/internal/store"
)

var usernames = []string{
	"Alice", "Bob", "Charlie", "Diana", "Eve", "Frank", "Grace", "Hank", "Ivy", "Jack",
	"Kara", "Leo", "Mona", "Nina", "Oscar", "Pia", "Quinn", "Rita", "Sam", "Tina",
	"Uma", "Vince", "Wade", "Xena", "Yara", "Zane", "Amber", "Ben", "Chloe", "Dexter",
	"Ella", "Finn", "Gwen", "Harry", "Iris", "Jade", "Kyle", "Luna", "Milo", "Nora",
	"Owen", "Paula", "Quincy", "Ralph", "Sophie", "Tom", "Ursula", "Vivian", "Walt", "Xander",
}

var titles = []string{
	"How to Start Your Day Right",
	"The Basics of Healthy Eating",
	"Mastering Time Management",
	"Top 10 Travel Destinations",
	"A Guide to Minimalist Living",
	"Fitness Tips for Beginners",
	"Why Reading Matters",
	"Exploring Local Cuisine",
	"The Art of Meditation",
	"Productivity Hacks You Need",
	"Tips for Better Sleep",
	"Creating a Home Office",
	"How to Budget Wisely",
	"Gardening for Beginners",
	"The Benefits of Journaling",
	"Learning a New Skill",
	"Eco-Friendly Lifestyle Tips",
	"Quick and Easy Recipes",
	"Building Positive Habits",
	"Finding Balance in Life",
}

var contents = []string{
	"Start your morning with a clear routine to set a positive tone for the day.",
	"Discover the essentials of balanced meals and their impact on your health.",
	"Learn time management strategies to stay productive and avoid burnout.",
	"Check out these beautiful places around the world for your next adventure.",
	"Simplify your life by focusing on what truly matters and decluttering your space.",
	"Beginner-friendly fitness tips to build a sustainable workout habit.",
	"Reading can expand your knowledge and improve your focus—here’s how to start.",
	"Explore must-try dishes that capture the flavors of local culture.",
	"Meditation can help reduce stress and improve mental clarity—here’s a guide.",
	"Use these tips to boost productivity and accomplish more with less effort.",
	"Unlock better sleep with these habits to enhance your nightly rest.",
	"Design a functional home office to improve work-from-home productivity.",
	"Tips on creating a budget that helps manage finances and grow savings.",
	"Gardening essentials for beginners to start their own garden at home.",
	"Journaling can improve mental health—here’s how to start your journey.",
	"Learn the basics to pick up new skills faster and more efficiently.",
	"Easy ways to adopt eco-friendly practices and reduce your carbon footprint.",
	"Try these quick, delicious recipes for busy weeknights.",
	"Practical strategies for building habits that stick and last long-term.",
	"Achieving a balanced lifestyle is essential for happiness and well-being.",
}

var tags = []string{
	"Health", "Travel", "Productivity", "Lifestyle", "Cooking",
	"Wellness", "Finance", "Technology", "Fitness", "Sustainability",
	"Mindfulness", "Self-Care", "Food", "Personal Development",
	"Outdoor", "Minimalism", "Work-Life Balance", "Hobbies",
	"Relationships", "Home Improvement",
}

var comments = []string{
	"Great insights! I’ll definitely try this out.",
	"Thanks for sharing! This was really helpful.",
	"I completely agree with your points here.",
	"This post inspired me to make some changes.",
	"Can you share more about this topic?",
	"I love the tips on productivity!",
	"Thanks for breaking it down so clearly.",
	"I learned something new today—thank you!",
	"Such a useful guide for beginners!",
	"I’ve been looking for advice like this.",
	"I appreciate your perspective on this.",
	"Can’t wait to read more from you!",
	"This really resonates with me.",
	"Great post! Keep up the good work.",
	"I never thought of it that way!",
	"The tips here are very practical.",
	"Your blog is a go-to resource for me.",
	"Interesting points, I’ll need to reflect on this.",
	"Thanks for the easy-to-follow steps.",
	"This post really helped me understand the basics.",
}

func Seed(store store.Storage, db *sql.DB) {
	ctx := context.Background()

	users := generateUsers(200)
	for _, user := range users {
		if err := store.Users.Create(ctx, user); err != nil {
			log.Println("error creating users", err)
			return
		}
	}

	posts := generatePosts(200, users)
	for _, post := range posts {
		if err := store.Posts.Create(ctx, post); err != nil {
			log.Println("error creating posts", err)
			return
		}
	}

	comments := generateComments(500, users, posts)
	for _, comment := range comments {
		if err := store.Comments.Create(ctx, comment); err != nil {
			log.Println("error creating posts", err)
			return
		}
	}

	log.Println("seeding is complete")
}

func generateUsers(num int) []*store.User {
	users := make([]*store.User, num)

	// Create a copy of usernames to avoid modifying the original slice
	availableUsernames := make([]string, len(usernames))
	copy(availableUsernames, usernames)

	// If we need more users than available usernames, we'll add numbers to make them unique
	for i := 0; i < num; i++ {
		var username string
		if i < len(availableUsernames) {
			username = availableUsernames[i]
		} else {
			// For additional users beyond the available usernames,
			// create unique names by adding random suffixes
			randomSuffix := rand.Intn(9999)
			baseUsername := availableUsernames[i%len(availableUsernames)]
			username = fmt.Sprintf("%s%d", baseUsername, randomSuffix)
		}

		// Create a unique email using a random number
		randomNum := rand.Intn(999999)
		email := fmt.Sprintf("%s.%d@myemail.com", strings.ToLower(username), randomNum)

		users[i] = &store.User{
			Username: username,
			Email:    email,
			Password: "123123",
		}
	}
	return users
}

func generatePosts(num int, users []*store.User) []*store.Post {
	posts := make([]*store.Post, num)

	for i := 0; i < num; i++ {
		user := users[rand.Intn(len(users))]

		posts[i] = &store.Post{
			UserID:  user.ID,
			Title:   titles[rand.Intn(len(titles))],
			Content: contents[rand.Intn(len(contents))],
			Tags: []string{
				tags[rand.Intn(len(tags))],
				tags[rand.Intn(len(tags))],
			},
		}
	}
	return posts
}

func generateComments(num int, users []*store.User, posts []*store.Post) []*store.Comment {
	cms := make([]*store.Comment, num)

	for i := 0; i < num; i++ {
		cms[i] = &store.Comment{
			PostID:  posts[rand.Intn(len(posts))].ID,
			UserID:  users[rand.Intn(len(users))].ID,
			Content: comments[rand.Intn(len(comments))],
		}
	}
	return cms
}
