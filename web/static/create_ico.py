#!/usr/bin/env python3
"""
Create a simple ICO file from scratch without external dependencies.
Creates a 32x32 hockey puck icon.
"""

import struct
import os

def create_simple_ico(output_path):
    """Create a simple 32x32 ICO with a hockey stick emoji design."""

    size = 32

    # Create a simple bitmap data (32x32, 32-bit RGBA)
    width, height = size, size

    # RGBA pixel data
    pixels = []
    center_x, center_y = width // 2, height // 2

    # Draw a simple hockey stick representation
    for y in range(height):
        for x in range(width):
            # White background
            r, g, b, a = 255, 255, 255, 255

            # Draw a stylized hockey stick (diagonal brown stick with black blade)
            # Stick shaft (diagonal from top-left to bottom-right)
            dx_stick = abs((x - y) - 2)
            if dx_stick < 3 and 4 < x < 24 and 4 < y < 24:
                # Brown stick color
                r, g, b, a = 139, 69, 19, 255
                # Add some shading
                if dx_stick < 2:
                    r, g, b = min(160, r + 20), min(90, g + 20), min(40, b + 20)

            # Blade at bottom (horizontal)
            if 20 < x < 28 and 22 < y < 28:
                # Black blade
                r, g, b, a = 30, 30, 30, 255
                # Highlight on blade
                if y < 25:
                    r, g, b = 50, 50, 50

            # Puck (small circle at bottom right)
            dx = x - 25
            dy = y - 20
            dist = (dx * dx + dy * dy) ** 0.5
            if dist < 2.5:
                r, g, b, a = 40, 40, 40, 255

            pixels.append((b, g, r, a))  # BMP is BGRA

    # Create ICO file
    # ICO header
    ico_header = struct.pack('<HHH', 0, 1, 1)  # Reserved, Type (1=ICO), Count

    # ICO directory entry
    bpp = 32
    image_size = width * height * 4  # RGBA
    colors = 0  # 0 means 256+ colors

    # Directory entry: width, height, colors, reserved, planes, bpp, size, offset
    ico_dir = struct.pack('<BBBBHHII',
                          width, height, colors, 0,  # width, height, colors, reserved
                          1, bpp,  # planes, bits per pixel
                          40 + image_size,  # size of image data (header + pixels)
                          22)  # offset (6 header + 16 directory)

    # BMP header (BITMAPINFOHEADER)
    bmp_header = struct.pack('<IIIHHIIIIII',
                            40,  # header size
                            width, height * 2,  # width, height*2 for ICO
                            1, bpp,  # planes, bits per pixel
                            0,  # compression (0 = none)
                            image_size,  # image size
                            0, 0, 0, 0)  # various unused fields

    # Combine pixel data (bottom-up for BMP)
    pixel_data = b''
    for y in range(height - 1, -1, -1):
        for x in range(width):
            idx = y * width + x
            b, g, r, a = pixels[idx]
            pixel_data += struct.pack('BBBB', b, g, r, a)

    # Write ICO file
    with open(output_path, 'wb') as f:
        f.write(ico_header)
        f.write(ico_dir)
        f.write(bmp_header)
        f.write(pixel_data)

    file_size = os.path.getsize(output_path)
    print(f"✓ Created {output_path}")
    print(f"  Size: {width}x{height}, {file_size} bytes")
    print(f"  Hockey stick design 🏒")

if __name__ == '__main__':
    import sys
    from pathlib import Path

    script_dir = Path(__file__).parent
    output = script_dir / 'favicon.ico'

    print("🏒 Creating hockey stick favicon.ico...")
    create_simple_ico(output)
    print("\n✅ Favicon created successfully!")
    print("   The icon features a hockey stick 🏒 with puck!")

