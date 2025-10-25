# -----------------------------
# Tiltfile for Go API Gateway
# -----------------------------
load('ext://git_resource', 'git_checkout')  # optional if you use Git dependencies

# Print a welcome message in Tilt UI
print("""
-----------------------------------------------------------------
✨ Tilt is running your Go API Gateway
-----------------------------------------------------------------
""")

# -----------------------------
# Run the Go API Gateway locally
# -----------------------------
local_resource(
    name='api-gateway',  # name of the resource in Tilt UI
    cmd='go run services/api-gateway/main.go',  # command to start the server
    deps=['services/api-gateway'],  # watch this folder for changes
    
)

# -----------------------------
# Optional: show instructions in Tilt UI
# -----------------------------
warn('ℹ️ API Gateway running locally. Any changes in services/api-gateway will auto-restart the server.')

# -----------------------------
# End of Tiltfile
# -----------------------------