import cv2

image = cv2.imread('gg.jpg', cv2.IMREAD_GRAYSCALE)
cv2.imwrite('g.jpg', image)