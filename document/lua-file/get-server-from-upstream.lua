local table = require("table")

local server_name = ngx.var.server_name;
ngx.log(ngx.ERR, " current server_name ", server_name);

local server_list_str = ngx.shared.backends_zone:get(server_name)
ngx.log(ngx.ERR, " current ups ", server_list_str)

if server_list_str == nil then
	ngx.log(ngx.ERR, "use default upstream");
	return "webapi_list";
end

if string.find(server_list_str, ",") == nil then
	return server_list_str
end

local arr = {}
for match in (server_list_str..","): gmatch("(.-)"..",") do
	table.insert(arr, match)
end
for i = 1, #arr do
	ngx.log(ngx.ERR, " arr has ", arr[i]);
end
local count = ngx.shared.backends_zone:get(server_name .. "-count")
if count == #arr then
	count = 1
elseif count == nil then
	count = 1
else
	count = count + 1
end
ngx.shared.backends_zone:set(server_name .. "-count", count);
ngx.log(ngx.ERR, " use list[] ", count, " server ", arr[count]);
return arr[count];