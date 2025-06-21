# Hands-on exercise #62 - interfaces

- create a type SQUARE
  - length float64
  - width float64
- create a type CIRCLE
  - radius float64
- attach a method to each that calculates AREA and returns it
  - circle area= π r 2
    - math.Pi
    - math.Pow
  - square area = L * W
- create a type SHAPE that defines an interface as anything that has the AREA method
- create a func INFO which takes type shape and then prints the area
- create a value of type square
- create a value of type circle
- use func info to print the area of square
- use func info to print the area of circle
