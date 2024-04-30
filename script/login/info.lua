local function getRuntimeOS()
    local OS = os.getenv("OS")

    if OS == "Windows_NT" then
        return "Windows"
    elseif os.execute("uname -s >/dev/null") then
        local uname = io.popen("uname -s"):read("*l")

        if uname == "Darwin" then
            return "macOS"
        end

        return uname
    end

    return "Unknown"
end

local function sendServerInfo(session)
    local go = getGoVersion()
    local os = getRuntimeOS()
    local build = getBuildInfo()

    sendClientMessage(session, 'Welcome to RXJH-EMU - Yulgang Server Emulator!')
    sendClientMessage(session,
        'Running on '..os..' OS with ' ..go..' and '.._VERSION)
    sendClientMessage(session, 'Build: #'..build)
end

addEventHandler('onPlayerLogin', function(session, name, state)
    -- fail to log in
    if not state then
        return
    end

    sendServerInfo(session)
end)