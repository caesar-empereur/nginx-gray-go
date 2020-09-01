
local table = require("table")

local server = ngx.req.get_uri_args()["server"]
local upstream = ngx.req.get_uri_args()["upstream"]

if server == nil then
	ngx.say(upstream, " server is null, return");
	return
end
ngx.log(ngx.ERR, "add start ------------------------------------------", upstream);

local server_list_str = ngx.shared.backends_zone:get(upstream)

if server_list_str == nil then
	server_list_str = server
	ngx.log(ngx.ERR, " current server_list_str is null, add ", server_list_str);
else
	ngx.log(ngx.ERR, " current server_list_str is not null ", server_list_str);
	server_list_str = server_list_str .. "," .. server;
end

ngx.shared.backends_zone:set(upstream, server_list_str);
ngx.say(upstream, " add success ", server_list_str)
