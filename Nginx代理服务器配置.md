配置正向HTTP代理
编辑 nginx.conf 文件:

114.114.114.114 国内;
8.8.8.8 国外

server {
    resolver 114.114.114.114;
    resolver_timeout 5s;
    listen 0.0.0.0:8088
    access_log  /var/log/nginx/access.proxy.log main;
    error_log   /var/log/nginx/error.proy.log;

    location / {
        proxy_pass http://$http_host$request_uri;
        proxy_set_header Host $http_host;
        proxy_buffers 256 4k;
        proxy_max_temp_file_size 0;
        proxy_connect_timeout 30;
        proxy_cache_valid 200 302 10m;
        proxy_cache_valid 301 1h;
        proxy_cache_valid any 1m;
    }
}

server {  
    resolver 114.114.114.114; 
    listen 8080;  
    location / {  
        proxy_pass http://$http_host$request_uri;
        proxy_set_header HOST $http_host;
        proxy_buffers 256 4k;
        proxy_max_temp_file_size 0k; 
        proxy_connect_timeout 30;
        proxy_send_timeout 60;
        proxy_read_timeout 60;
        proxy_next_upstream error timeout invalid_header http_502;
    }  
}



重启nginx:
./sbin/nginx -s reload

配置
linux
export http_proxy=118.210.42.251:8808
export https_proxy=118.210.42.251:8808
要取消该设置：

unset http_proxy
unset https_proxy
windows
浏览器配置HTTP代理，配置完成

注意
nginx正向代理不支持代理 Https 网站
因为 Nginx 不支持 CONNECT，所以无法正向代理 Https 网站（网上银行，Gmail）。
如果访问 Https 网站，比如：https://www.google.com，Nginx access.log 日志如下：

“CONNECT www.google.com:443 HTTP/1.1” 400
