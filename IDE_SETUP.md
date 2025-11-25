# IDE Run Configurations Setup Complete! ✅

## What Was Created

I've set up **6 run configurations** for JetBrains GoLand/IntelliJ IDEA to make development easier:

### 🚀 Server Configurations

1. **Run Server** 
   - Port: 8080
   - Quick start for development
   - Most commonly used

2. **Run Server (Port 3000)**
   - Port: 3000
   - Use when 8080 is already in use

3. **Run Server (with build info)**
   - Port: 8080
   - Includes version/build metadata via ldflags
   - Similar to `make build`

4. **Run Server (with tests)**
   - Port: 8080
   - Runs all tests before starting the server
   - Ensures code quality before running

### 🧪 Test Configurations

5. **Test All**
   - Runs all tests in the project
   - Equivalent to `go test ./...`

6. **Test Player Search**
   - Runs only player search feature tests
   - Tests: `TestSearchPlayers`, `TestGetPlayerCareer`

## 📁 Files Created

```
.idea/runConfigurations/
├── README.md                                 # Detailed documentation
├── Run_Server.xml                            # Main server config
├── Run_Server__Port_3000_.xml               # Alternative port
├── Run_Server__with_build_info_.xml         # With build metadata
├── Run_Server__with_tests_.xml              # Run tests first
├── Test_All.xml                              # All tests
└── Test_Player_Search.xml                    # Player search tests
```

## 🎯 How to Use

### In GoLand/IntelliJ IDEA:

1. **Open the project** in your IDE (if not already open)

2. **Look at the top-right toolbar** - you'll see a dropdown with run configurations

3. **Select a configuration**:
   - "Run Server" - for quick development
   - "Run Server (with tests)" - to ensure tests pass first
   - "Test All" - to run all tests

4. **Click the green Run button** (▶) or press:
   - `Shift+F10` (Windows/Linux)
   - `Ctrl+R` (Mac)

5. **To debug**, click the Debug button (🐛) or press:
   - `Shift+F9` (Windows/Linux)
   - `Ctrl+D` (Mac)

### Quick Actions:

- **Run current file**: `Ctrl+Shift+F10` (Windows/Linux) or `Ctrl+Shift+R` (Mac)
- **Show run menu**: `Alt+Shift+F10` (Windows/Linux) or `Ctrl+Alt+R` (Mac)
- **Stop running process**: `Ctrl+F2` (Windows/Linux) or `Cmd+F2` (Mac)

## 🌐 Access the Application

After starting the server:

| Feature | URL |
|---------|-----|
| Team Roster | http://localhost:8080/roster |
| Player Career Search | http://localhost:8080/player-search |
| Build Info | http://localhost:8080/build-info |

## 🔧 Customize Configurations

To modify any configuration:

1. Click the run configuration dropdown
2. Select "Edit Configurations..."
3. Choose the configuration to edit
4. Modify settings (port, environment variables, etc.)
5. Click "Apply" and "OK"

### Common Customizations:

- **Change port**: Edit the `PORT` environment variable
- **Add environment variables**: In "Environment" section
- **Change working directory**: Modify "Working directory" field
- **Add program arguments**: In "Program arguments" field

## 💡 Tips

### Using Multiple Configurations

You can run multiple configurations simultaneously:
- Run the server on port 8080
- Run tests in parallel
- Each gets its own console tab

### Keyboard Shortcuts

- `Ctrl+Shift+A` (Windows/Linux) or `Cmd+Shift+A` (Mac) - Search for any action
- Type "Run" to see all run-related commands

### Debugging

All server configurations support debugging:
1. Set breakpoints by clicking in the left margin of code
2. Click the Debug button instead of Run
3. Use the debugger toolbar to step through code

## 🐛 Troubleshooting

**Configurations not appearing?**
- Restart the IDE
- Check that `.idea/runConfigurations/` directory exists
- Verify XML files are present

**Port already in use?**
- Use "Run Server (Port 3000)" configuration
- Or stop existing process: `pkill -f "target/bin/server"`

**Tests failing?**
- Check internet connection (tests call NHL API)
- Run `go mod tidy` to update dependencies

## ✨ Benefits

Using these run configurations instead of terminal:

✅ **One-click start** - No need to type commands  
✅ **Integrated console** - Output appears in IDE  
✅ **Easy debugging** - Set breakpoints and step through code  
✅ **Environment management** - Easy to change ports/variables  
✅ **Stop with one click** - No need to find and kill processes  
✅ **Multiple runs** - Run server and tests simultaneously  

## 📚 Additional Resources

- [GoLand Run Configurations](https://www.jetbrains.com/help/go/run-debug-configuration.html)
- [Debugging in GoLand](https://www.jetbrains.com/help/go/debugging-code.html)
- Project README: `README.md`
- Feature documentation: `FEATURE_SUMMARY.md`

---

**Happy coding! 🚀**

The server is now ready to run with a single click in your IDE!

