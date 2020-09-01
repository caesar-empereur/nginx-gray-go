
local server = ngx.req.get_uri_args()["server"]
local upstream = ngx.req.get_uri_args()["upstream"]
if server == nil then
	ngx.say(host, " server is null, return");
	return
end
ngx.log(ngx.ERR, "remove start ------------------------------------------", upstream);

local server_list_str = ngx.shared.backends_zone:get(upstream)
ngx.log(ngx.ERR, "server_list_str current ", server_list_str);

if server_list_str == nil then
	ngx.say(upstream, " current ups is null, return ")
	return
end

if string.find(server_list_str, server..",") ~= nil then
	server_list_str = string.gsub(server_list_str, server..",", "")
elseif string.find(server_list_str, ","..server) ~= nil then
	server_list_str = string.gsub(server_list_str, ","..server, "")
elseif string.find(server_list_str, server) ~= nil then
	server_list_str = string.gsub(server_list_str, server, "")
else
	ngx.log(ngx.ERR, " server not in list ", server);
end
ngx.say(upstream, " server_list_str after remove ", server_list_str)
ngx.shared.backends_zone:set(upstream, server_list_str);
ngx.say(upstream, " remove success ", server)
