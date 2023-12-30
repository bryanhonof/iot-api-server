from dataclasses import dataclass

@dataclass
class RGBdata:
    RED: int
    GREEN: int
    BLUE: int
    BRIGHTNESS: int
    

    def __init__(self, red: int = 0, green: int = 0, blue: int = 0, brightness: int =0):
        def limit_to_8bits(value):
            return value & 0xFF

        self.RED = limit_to_8bits(red)
        self.BLUE = limit_to_8bits(blue)
        self.GREEN = limit_to_8bits(green)
        self.BRIGHTNESS = limit_to_8bits(brightness)

    def output_bytes(self) -> bytes:

        # Calculate RGB values with brightness adjustment
        red_value = int(self.RED * self.BRIGHTNESS/255 )
        green_value = int(self.GREEN * self.BRIGHTNESS/255)
        blue_value = int(self.BLUE * self.BRIGHTNESS/255)

        # Pack RGB values into a 3-byte representation (bytes)
        output = bytes([blue_value, red_value, green_value])
        return output

    def colors(self):
        output: str = " "+str(self.RED) +" "+ str(self.GREEN) +" "+ str(self.BLUE) +" "+ str(self.BRIGHTNESS)
        return output
    
@dataclass
class Neopixel:
    
    def __init__(self, count: int =0):
        self.pixels = [RGBdata()] * count
        
    pixels = []
    
    def colors(self):
        output_list = []
        index = 0
        for pixel in self.pixels:
            output_list.append("LED" + str(index) + self.pixels[index].colors())
            index += 1
        return ' '.join(output_list)


    def set_pixel(self,pixel: int , input: RGBdata):
        self.pixels[pixel] = input
        return 0

    def fill(self, input: RGBdata):
        self.pixels = [input] * len(self.pixels)
        return 0

    def ws2812_Data(self):
        outputArray = [0] * (len(self.pixels) * 24)
        index = 0

        for pixel in self.pixels:
            color = pixel.output_bytes()

            for i in range(23, -1, -1):
                byte_index = i // 8
                bit_index = 7 - (i % 8)

                if ((color[byte_index] >> bit_index) & 0x01) == 1:
                    outputArray[index] = 0b11111000  # store 1
                else:
                    outputArray[index] = 0b11100000 # store 0
                index += 1

        return outputArray