#!/usr/bin/env python3
"""
Generate favicon.ico from SVG source.
Requires: pip install pillow cairosvg
"""

import io
from pathlib import Path

try:
    from PIL import Image
    import cairosvg
except ImportError:
    print("Installing required packages...")
    import subprocess
    import sys
    subprocess.check_call([sys.executable, "-m", "pip", "install", "pillow", "cairosvg"])
    from PIL import Image
    import cairosvg

def svg_to_ico(svg_path, ico_path):
    """Convert SVG to multi-resolution ICO file."""

    # Read SVG
    with open(svg_path, 'r') as f:
        svg_data = f.read()

    # Sizes for favicon (common sizes)
    sizes = [16, 32, 48, 64]
    images = []

    for size in sizes:
        # Convert SVG to PNG at specific size
        png_data = cairosvg.svg2png(
            bytestring=svg_data.encode('utf-8'),
            output_width=size,
            output_height=size
        )

        # Load as PIL Image
        img = Image.open(io.BytesIO(png_data))
        images.append(img)
        print(f"✓ Generated {size}x{size} icon")

    # Save as ICO with multiple resolutions
    images[0].save(
        ico_path,
        format='ICO',
        sizes=[(img.width, img.height) for img in images],
        append_images=images[1:]
    )

    print(f"\n✓ Created {ico_path}")
    print(f"  File size: {Path(ico_path).stat().st_size} bytes")

if __name__ == '__main__':
    script_dir = Path(__file__).parent
    svg_path = script_dir / 'favicon.svg'
    ico_path = script_dir / 'favicon.ico'

    print("Generating favicon.ico from SVG...")
    print(f"Source: {svg_path}")
    print(f"Output: {ico_path}\n")

    svg_to_ico(svg_path, ico_path)

    print("\n🎉 Favicon generated successfully!")
    print("The favicon includes 16x16, 32x32, 48x48, and 64x64 resolutions.")

