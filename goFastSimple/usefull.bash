apt-get autoremove --purge apt-xapian-index   !! We dont neet it?!
apt-get autoremove --purge 

 echo 524288 > /proc/sys/fs/inotify/max_user_watches
 vi /etc/sysctl.conf
 fs.inotify.max_user_watches=1900000
 vm.overcommit_memory = 1  ??
 sysctl -p


Check nginx 
strace -p 3544 -p 3545 -p 3546 -p 3547 2>&1 | grep gz

ps -eo %mem,rss,pid,args --sort rss | tail -10

dd if=/dev/zero of=/swapfile bs=1024 count=128k

mkswap /swapfile
swapon /swapfile
swapon -s
vi /etc/fstab
/swapfile       none    swap    sw      0       0 

vm.swappiness = 10  in /etc/sysctl.conf  "sysctl -p"
cat /proc/sys/vm/swappiness


GOPATH=$GOPATH:/home/juno/git/goFastSimple/goFastSimple go test -v


bin/syncpushdomaindb -locale=fi_FI -themes=porno

ZRANGEBYSCORE pushdomains -inf +inf withscores  LIMIT 0 10000
ZRANGEBYSCORE limitsites -inf +inf withscores  LIMIT 0 10000
ZRANGEBYSCORE fi_FI:porno:keywords -inf +inf withscores  LIMIT 0 10000

SRANDMEMBER pagetocreate

ZADD pushdomains 1 pilluseksi.com
ZADD pushdomains 1 pilluseksi.info


stop gofast && start gofast && stop nginx && start nginx
grep -E 'www.google|bing'  logs/access.log |grep -E 'url|search'

cat logs/access.log |grep -E 'www.google|bing' |grep q= |grep -v .js |grep -v .css|grep -v .woff|grep -v .eot  |grep -E 'url|search'


dpkg-reconfigure tzdata

echo 'export TZ="Europe/Helsinki"' >> /etc/default/rsyslog
restart rsyslog
adduser juno
passwd

apt-get install git-core
apt-get install redis-server

apt-get install libgeoip1 libgeoip-dev  ??


SELECT * FROM keywords where themes='porno' and Hits=0 and Updated <1410838360865
SELECT * FROM keywords where themes='porno' and Hits=1 and Updated <

select Phrase from phrases  group by Phrase having count(*) >1




scp /home/juno/git/goFastCgi/goFastCgi/singo.db 104.131.209.134:git/goFastCgiLight/goFastCgiLight
scp /home/juno/git/goFastCgi/goFastCgi/singo.db 104.131.99.251:git/goFastCgiLight/goFastCgiLight

bin/startsite -locale=fi_FI -themes=porno -site= -variant=0


grep -E 'www.google|bing' fi_FIporno_access.log |grep -E 'url|search'
grep -E 'www.google|bing' logs/access.log |grep -E 'url|search'

touch -t 201101250000 dummy
touch dummy
find www -newer dummy -exec ls -l  {} \;

# if was done some local modification
git stash save --keep-index
git stash drop


