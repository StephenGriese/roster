#!/bin/bash
# Generate favicon from SVG using macOS built-in tools and optionally sharp/node

set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
SVG_FILE="$SCRIPT_DIR/favicon.svg"
ICO_FILE="$SCRIPT_DIR/favicon.ico"

echo "🏒 Generating hockey-themed favicon..."
echo "Source: $SVG_FILE"
echo ""

# Check if we have node and can use sharp
if command -v node &> /dev/null; then
    echo "Using Node.js with sharp for high-quality conversion..."

    # Create a temporary Node.js script
    cat > /tmp/gen_favicon.js << 'EOJS'
const fs = require('fs');
const path = require('path');

// Try to use sharp if available, otherwise provide instructions
try {
    const sharp = require('sharp');

    const svgPath = process.argv[2];
    const icoPath = process.argv[3];

    console.log('Converting SVG to multi-resolution ICO...');

    // Create 32x32 favicon (most common size)
    sharp(svgPath)
        .resize(32, 32)
        .png()
        .toFile(icoPath.replace('.ico', '_32.png'))
        .then(() => {
            console.log('✓ Generated 32x32 PNG');
            return sharp(svgPath).resize(16, 16).png().toFile(icoPath.replace('.ico', '_16.png'));
        })
        .then(() => {
            console.log('✓ Generated 16x16 PNG');
            console.log('\n✅ PNG files created!');
            console.log('Note: For full ICO support, you can:');
            console.log('1. Use an online converter: https://convertio.co/png-ico/');
            console.log('2. Or install ImageMagick: brew install imagemagick');
            console.log('   Then run: convert favicon_16.png favicon_32.png favicon.ico');
        })
        .catch(err => {
            console.error('Error:', err);
            process.exit(1);
        });

} catch(e) {
    console.log('Sharp not found. Installing...');
    console.log('Run: npm install sharp');
    console.log('Or use the SVG directly in your HTML.');
    process.exit(1);
}
EOJS

    node /tmp/gen_favicon.js "$SVG_FILE" "$ICO_FILE" || {
        echo ""
        echo "⚠️  Sharp not installed. Installing sharp..."
        cd "$SCRIPT_DIR"
        npm install --no-save sharp 2>/dev/null || {
            echo ""
            echo "📝 Manual conversion steps:"
            echo "1. Visit: https://convertio.co/svg-ico/"
            echo "2. Upload: $SVG_FILE"
            echo "3. Download as favicon.ico"
            echo ""
            echo "Or install ImageMagick: brew install imagemagick"
            echo "Then run: convert favicon.svg -define icon:auto-resize=64,48,32,16 favicon.ico"
        }
    }
else
    echo "Node.js not found."
    echo ""
    echo "📝 To convert the SVG to ICO, use one of these methods:"
    echo ""
    echo "Method 1: Online converter (easiest)"
    echo "  1. Visit: https://convertio.co/svg-ico/"
    echo "  2. Upload: $SVG_FILE"
    echo "  3. Download and save as: $ICO_FILE"
    echo ""
    echo "Method 2: ImageMagick (if installed)"
    echo "  brew install imagemagick"
    echo "  convert favicon.svg -define icon:auto-resize=64,48,32,16 favicon.ico"
    echo ""
    echo "Method 3: Use SVG directly (modern browsers)"
    echo "  Add to HTML: <link rel=\"icon\" type=\"image/svg+xml\" href=\"/favicon.svg\">"
fi

echo ""
echo "✨ SVG favicon created at: $SVG_FILE"
echo "   You can use this directly in modern browsers!"

