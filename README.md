# VideoProcessing 

## Technology

- Go
- gocv (OpenCV for Golang)

## How to run

``` go run main.go ```


## Commands

<table>
  <tr>
   <td><strong>Tecla</strong>
   </td>
   <td><strong>Efeito</strong>
   </td>
   <td><strong>Descrição</strong>
   </td>
  </tr>
  <tr>
   <td>1
   </td>
   <td>GaussianBlur
   </td>
   <td>Trackbar level controla o tamanho do kernel, determinando o nível de "borramento"
<p>
Min=1 Max=201
   </td>
  </tr>
  <tr>
   <td>2
   </td>
   <td>Canny
   </td>
   <td>Trackbar level controla o threshold 1.
<p>
Min=0 Max=300
   </td>
  </tr>
  <tr>
   <td>3
   </td>
   <td>Sobel
   </td>
   <td>
   </td>
  </tr>
  <tr>
   <td>4
   </td>
   <td>Grayscale
   </td>
   <td>
   </td>
  </tr>
  <tr>
   <td>5
   </td>
   <td>Brightness
   </td>
   <td>Trackbar level controla o brilho da imagem
<p>
Min=-50 Max=50
   </td>
  </tr>
  <tr>
   <td>6
   </td>
   <td>Contrast
   </td>
   <td>Trackbar level controla o contraste. \
Min=0 Max=10
   </td>
  </tr>
  <tr>
   <td>7
   </td>
   <td>Resize
   </td>
   <td>Trackbar level controla o tamanho da imagem resultante.
<p>
Min=(Cols,Rows)
<p>
Max=(Cols,Rows) / 26
   </td>
  </tr>
  <tr>
   <td>V
   </td>
   <td>VerticalFlip
   </td>
   <td>
   </td>
  </tr>
  <tr>
   <td>H
   </td>
   <td>HorizontalFlip
   </td>
   <td>
   </td>
  </tr>
  <tr>
   <td>ARROW LEFT
   </td>
   <td>RotateCounterClockwise
   </td>
   <td>
   </td>
  </tr>
  <tr>
   <td>ARROW RIGHT
   </td>
   <td>RotateClockwise
   </td>
   <td>
   </td>
  </tr>
  <tr>
   <td>ARROW UP
   </td>
   <td>Rotate 180
   </td>
   <td>
   </td>
  </tr>
  <tr>
   <td>R
   </td>
   <td>Gravação
   </td>
   <td>Grava um vídeo com o efeito resultante no arquivo "video.mov"
   </td>
  </tr>
</table>
