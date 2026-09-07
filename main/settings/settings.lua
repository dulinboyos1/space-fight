-- This file is responsible for reading the settings file into memory and sharing it with all files
-- The loaded data is cached by Lua natively so only 1 file read will happen per boot (as intended)

-- Files should not make copies of the settings table that way when a change is made it has immediate effect

local CLEAN = true -- Clear next launch

if CLEAN then
	print("********* WARNING ****************** WARNING ****************** WARNING *********")
	print("Save file clean directive is set: Settings file is being reset to defaults")
	print("********* WARNING ****************** WARNING ****************** WARNING *********")
end

local path = sys.get_save_file("space-fight", "settings")
local settings = sys.load(path)

if next(settings) == nil or CLEAN then -- Empty table, needs initalizing
	print("No settings data found initalizing file with defaults")
	settings = json.decode(sys.load_resource("/data/settings.json")) -- Defaults
	sys.save(path, settings) -- Init as default
end

-- Blank the subscribers ready for use
settings.subscribers = {}

return settings