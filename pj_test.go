package main

import (
	"fmt"
	"os"
	"testing"
	"github.com/joho/godotenv"
)

func Test_Dependency_score(t *testing.T) {
	fmt.Println("TEST 1")
	owner := "cloudinary"
	repo := "cloudinary_npm"
	score := dependency_score(owner, repo)
	fmt.Println(score)
}

func Test_ramp_up_score1(t *testing.T) {
	fmt.Println("TEST 2")
	url := "https://github.com/lodash/lodash"
	score, owner, repo := ramp_up_score(url)
	fmt.Printf("%f\n", score)
	fmt.Println(owner)
	fmt.Println(repo)
}

func Test_ramp_up_score2(t *testing.T) {
	fmt.Println("TEST 3")
	url := "https://github.com/nullivex/nodist"
	score, owner, repo := ramp_up_score(url)
	fmt.Printf("%f\n", score)
	fmt.Println(owner)
	fmt.Println(repo)
}

func Test_ramp_up_score3(t *testing.T) {
	fmt.Println("TEST 4")
	url := "https://github.com/cloudinary/cloudinary_npm"
	score, owner, repo := ramp_up_score(url)
	fmt.Printf("%f\n", score)
	fmt.Println(owner)
	fmt.Println(repo)
}

func Test_correctness_score1(t *testing.T) {
	fmt.Println("TEST 5")
	var personal_token string
	personal_token = os.Getenv("GITHUB_TOKEN")
	owner := "cloudinary"
	repo := "cloudinary_npm"
	score := correctness_score(personal_token, owner, repo)
	fmt.Printf("%f\n", score)
}

func Test_correctness_score2(t *testing.T) {
	fmt.Println("TEST 6")
	var personal_token string
	personal_token = os.Getenv("GITHUB_TOKEN")
	owner := "lodash"
	repo := "lodash"
	score := correctness_score(personal_token, owner, repo)
	fmt.Printf("%f\n", score)
}

func Test_correctness_score3(t *testing.T) {
	fmt.Println("TEST 7")
	var personal_token string
	personal_token = os.Getenv("GITHUB_TOKEN")
	owner := "nullivex"
	repo := "nodist"
	score := correctness_score(personal_token, owner, repo)
	fmt.Printf("%f\n", score)
}

func Test_responseviness_score1(t *testing.T) {
	fmt.Println("TEST 8")
	var personal_token string
	personal_token = os.Getenv("GITHUB_TOKEN")
	owner := "xingyizhou"
	repo := "CenterNet"
	score := responseviness_score(personal_token, owner, repo)
	fmt.Printf("%f\n", score)
}

func Test_responseviness_score2(t *testing.T) {
	fmt.Println("TEST 9")
	var personal_token string
	personal_token = os.Getenv("GITHUB_TOKEN")
	owner := "hwholiday"
	repo := "short_url"
	responseviness_score(personal_token, owner, repo)
}

func Test_responseviness_score3(t *testing.T) {
	fmt.Println("TEST 10")
	var personal_token string
	personal_token = os.Getenv("GITHUB_TOKEN")
	owner := "Caturra000"
	repo := "co"
	responseviness_score(personal_token, owner, repo)
}

func Test_bus_factor_score1(t *testing.T) {
	fmt.Println("TEST 11")
	var personal_token string
	personal_token = os.Getenv("GITHUB_TOKEN")
	owner := "cloudinary"
	repo := "cloudinary_npm"
	score := bus_factor_score(personal_token, owner, repo)
	fmt.Printf("%f\n", score)
}

func Test_bus_factor_score2(t *testing.T) {
	fmt.Println("TEST 12")
	var personal_token string
	personal_token = os.Getenv("GITHUB_TOKEN")
	owner := "lodash"
	repo := "lodash"
	score := bus_factor_score(personal_token, owner, repo)
	fmt.Printf("%f\n", score)
}

func Test_bus_factor_score3(t *testing.T) {
	fmt.Println("TEST 13")
	var personal_token string
	personal_token = os.Getenv("GITHUB_TOKEN")
	owner := "nullivex"
	repo := "nodist"
	score := bus_factor_score(personal_token, owner, repo)
	fmt.Printf("%f\n", score)
}

func Test_license_score1(t *testing.T) {
	fmt.Println("TEST 14")
	var personal_token string
	personal_token = os.Getenv("GITHUB_TOKEN")
	owner := "xingyizhou"
	repo := "CenterNet"
	score := license_score(personal_token, owner, repo)
	fmt.Printf("%f\n", score)
}

func Test_license_score2(t *testing.T) {
	fmt.Println("TEST 15")
	var personal_token string
	personal_token = os.Getenv("GITHUB_TOKEN")
	owner := "hwholiday"
	repo := "short_url"
	score := license_score(personal_token, owner, repo)
	fmt.Printf("%f\n", score)
}

func Test_license_score3(t *testing.T) {
	fmt.Println("TEST 16")
	var personal_token string
	personal_token = os.Getenv("GITHUB_TOKEN")
	owner := "Caturra000"
	repo := "co"
	score := license_score(personal_token, owner, repo)
	fmt.Printf("%f\n", score)
}

func Test_get_git_url_empty(t *testing.T) {
	fmt.Println("TEST 17")
	get_git_url("")
}

func Test_get_git_url1(t *testing.T) {
	fmt.Println("TEST 18")
	url := get_git_url("https://www.npmjs.com/package/express")
	fmt.Println(url)
}

func Test_analyze_git_empty(t *testing.T) {
	fmt.Println("TEST 19")
	analyze_git("", "")
}

func Test_analyze_git1(t *testing.T) {
	fmt.Println("TEST 20")
	analyze_git("https://github.com/cloudinary/cloudinary_npm", "https://github.com/cloudinary/cloudinary_npm")
}

func Test_entire_program(t *testing.T) {
	fmt.Println("TEST 21")
	calc_score("./sample_url_file.txt")
}

func Test_code_review_metric(t *testing.T) {
	fmt.Println("TEST 22")
	var personal_token string
	personal_token = os.Getenv("GITHUB_TOKEN")
	owner := "lodash"
	repo := "lodash"
	score := code_review_metric(personal_token, owner, repo)
	fmt.Printf("%f\n", score)
}

func Test_code_review_metric_2(t *testing.T) {
	fmt.Println("TEST 23")
	var personal_token string
	personal_token = os.Getenv("GITHUB_TOKEN")
	owner := "nullivex"
	repo := "nodist"
	score := code_review_metric(personal_token, owner, repo)
	fmt.Printf("%f\n", score)
}

func Test_code_review_metric_3(t *testing.T) {
	fmt.Println("TEST 24")
	godotenv.Load()
	var personal_token string
	personal_token = os.Getenv("GITHUB_TOKEN")
	owner := "cloudinary"
	repo := "cloudinary_npm"
	score := code_review_metric(personal_token, owner, repo)
	fmt.Printf("%f\n", score)
}
