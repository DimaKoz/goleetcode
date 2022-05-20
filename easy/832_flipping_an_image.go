package easy

/*
https://leetcode.com/problems/flipping-an-image/

Constraints:

n == image.length
n == image[i].length
1 <= n <= 20
images[i][j] is either 0 or 1.
*/

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Flipping an Image.
Memory Usage: 2.9 MB, less than 74.32% of Go online submissions for Flipping an Image.
*/

func flipAndInvertImage(image [][]int) [][]int {
	endindex := 0
	for i := 0; i < len(image); i++ {
		//flipping
		for ii := len(image[i]) - 1; ii > (len(image[i])-1)/2; ii-- {
			endindex = len(image[i]) - ii - 1
			image[i][ii], image[i][endindex] = image[i][endindex], image[i][ii]
		}
		//inverting
		for ii := 0; ii < len(image[i]); ii++ {
			if image[i][ii] == 0 {
				image[i][ii] = 1
			} else {
				image[i][ii] = 0
			}
		}
	}
	return image
}
