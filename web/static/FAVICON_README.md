# Favicon Documentation

## Overview

The application now has a professional hockey-themed favicon featuring a 3D hockey puck design.

## Files

### favicon.svg
- **Format**: SVG (Scalable Vector Graphics)
- **Size**: ~1.2 KB
- **Design**: 3D hockey puck with NHL-style colors
  - Ice blue background (#0051a5)
  - Black/grey puck with 3D shading
  - Red NHL-style stripe (#c8102e)
  - "R" monogram for Roster
- **Browser Support**: All modern browsers (Chrome, Firefox, Safari, Edge)
- **Advantage**: Scales perfectly at any size, small file size

### favicon.ico
- **Format**: ICO (Windows Icon)
- **Size**: ~4.2 KB
- **Resolution**: 32x32 pixels, 32-bit color
- **Design**: Hockey puck with NHL colors
- **Browser Support**: All browsers including older versions
- **Purpose**: Fallback for older browsers and Windows taskbar/bookmarks

## Implementation

The favicons are referenced in the HTML `<head>` section (see `server/components.go`):

```go
h.Link(h.Rel("icon"), h.Type("image/svg+xml"), h.Href("/favicon.svg")),
h.Link(h.Rel("alternate icon"), h.Href("/favicon.ico")),
```

Modern browsers will use the SVG version for crisp display at any size. Older browsers will fall back to the ICO version.

## Design Elements

The favicon features:

1. **Hockey Puck**: Central element representing the sport
2. **3D Effect**: Shading and highlights for depth
3. **NHL Colors**: 
   - Blue (#0051a5): Official NHL blue
   - Red (#c8102e): NHL red accent stripe
4. **Professional Look**: Clean, recognizable at small sizes
5. **Brand Element**: Subtle "R" for "Roster"

## Updating the Favicon

### Option 1: Modify the SVG
Edit `favicon.svg` with any text editor or vector graphics program (like Inkscape, Adobe Illustrator, or Figma).

### Option 2: Generate New ICO
If you update the SVG, regenerate the ICO file:

```bash
cd web/static
python3 create_ico.py
```

This creates a new `favicon.ico` from the design programmatically.

### Option 3: Use Online Tools
1. Edit `favicon.svg` as desired
2. Visit https://convertio.co/svg-ico/
3. Upload your SVG
4. Download as `favicon.ico`

## Browser Display

The favicon appears in:
- Browser tabs
- Bookmarks/favorites
- Browser history
- Desktop shortcuts (when pinning sites)
- Mobile home screen (when saved as web app)

## Technical Details

### SVG Specifications
- ViewBox: 64x64
- Color Space: RGB
- Shape: Rounded rectangle background (border-radius: 8px)
- Elements: Ellipses, rectangles, text

### ICO Specifications
- Format: Microsoft ICO
- Color Depth: 32-bit RGBA
- Dimensions: 32x32 pixels
- File Structure: ICO header + directory + BMP image data

## Troubleshooting

### Favicon Not Showing?
1. **Clear browser cache**: Hard refresh with `Cmd+Shift+R` (Mac) or `Ctrl+Shift+R` (Windows)
2. **Check file serving**: Visit http://localhost:8080/favicon.svg directly
3. **Restart server**: Stop and restart the application
4. **Check browser console**: Look for 404 errors

### Wrong Icon Showing?
- Browsers cache favicons aggressively
- Try in incognito/private mode
- Clear browser favicon cache specifically
- Wait a few minutes for cache to expire

## Generator Scripts

### create_ico.py
Pure Python script that creates the ICO file without external dependencies.
- Creates 32x32 ICO with hockey puck design
- Uses only Python standard library
- Generates proper BMP-based ICO format

### generate_favicon.py (Advanced)
Requires `pillow` and `cairosvg` for multi-resolution ICO creation.
- Creates 16x16, 32x32, 48x48, 64x64 icons
- Better quality than manual creation
- Run: `pip install pillow cairosvg && python3 generate_favicon.py`

### generate_favicon.sh
Shell script that attempts multiple conversion methods.
- Tries Node.js with sharp
- Falls back to manual instructions
- Provides multiple options

## Credits

Design inspired by:
- NHL official branding
- Modern flat design principles
- Hockey equipment aesthetics

Created for the NHL Roster application.
Last updated: November 25, 2025

