# Go言語の勉強

## # 環境構築

### ## 導入

以下を順に実行する．
保険のため、mainブランチへのpushはできないようにする．
```
$ git clone https://github.com/Knpand/Go-Study.git
$ cd Go-Study
$ docker-compose build
$ docker-compose up -d
```

### ## 環境

```
frontend:Vue.js (port=80)
backend:Go (port=5050)
```

### ## 目標
* CORSの理解を進める  
* いつもsessionの認証だからJWT認証サーバたててみる  
* 可能な限りサードパーティ製のパッケージ使わない  

#### 勉強内容ごとにbranch切る
現状はfrontからbackの5050/taskaにリクエスト飛ばしてもCORSではじかれる
