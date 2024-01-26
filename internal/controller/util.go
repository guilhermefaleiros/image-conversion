package controller

var ValidEffects = []string{"jpeg", "grayscale", "sepia", "png", "invert_colors"}

func Contains(slice []string, str string) bool {
	for _, item := range slice {
		if item == str {
			return true
		}
	}
	return false
}
