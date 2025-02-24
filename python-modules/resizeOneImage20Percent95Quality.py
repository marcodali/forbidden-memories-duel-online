from PIL import Image
import sys
from pathlib import Path

SCALE = 0.2 # this means resize to 20% its original size
QUALITY = 95

def process_image(input_path: str) -> None:
    with Image.open(input_path) as img:
        if img.mode in ('RGBA', 'P'):
            img = img.convert('RGB')
            
        width, height = img.size
        data = list(img.getdata())
        img_without_exif = Image.new(img.mode, img.size)
        img_without_exif.putdata(data)
        
        scale = SCALE
        new_width = int(width * scale)
        new_height = int(height * scale)
        
        resized_img = img_without_exif.resize((new_width, new_height), Image.Resampling.LANCZOS)
        
        output_path = str(Path(input_path).with_stem(f"{Path(input_path).stem}_copy"))
        resized_img.save(
            output_path,
            'JPEG',
            quality=QUALITY,
            optimize=True
        )

def main():
    if len(sys.argv) != 2:
        sys.exit(1)
    process_image(sys.argv[1])

if __name__ == "__main__":
    main()