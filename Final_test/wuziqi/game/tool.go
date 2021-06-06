package game

//判断胜利
func (r *Room)Judge(x,y int)bool{
	//横
	count:=0
	for i:=x-1;i>=0;i--{
		if r.Chess[x][y]==r.Chess[i][y]{
			count++
		}else{
			break
		}
	}
	for i:=x+1;i<=9;i++{
		if r.Chess[x][y]==r.Chess[i][y]{
			count++
		}else{
			break
		}
	}
	if count>=5{
		return true
	}
	//竖
	count=0
	for i:=y-1;i>=0;i--{
		if r.Chess[x][y]==r.Chess[x][i]{
			count++
		}else{
			break
		}
	}
	for i:=y+1;i<=9;i++{
		if r.Chess[x][y]==r.Chess[x][i]{
			count++
		}else{
			break
		}
	}
	if count>=5{
		return true
	}
	//右上
	count=0
	for i,j:=x-1,y-1;i>=0&&j>=0;i,j=i-1,j-1 {
		if r.Chess[x][y]==r.Chess[i][j]{
			count++
		}else{
			break
		}
	}
	for i,j:=x+1,y+1;i<=9&&j<=9;i,j=i+1,j+1{
		if r.Chess[x][y]==r.Chess[i][j]{
			count++
		}else{
			break
		}
	}
	if count>=5{
		return true
	}
	//右下
	count=0
	for i,j:=x-1,y+1;i>=0&&j<=9;i,j=i-1,j+1{
		if r.Chess[x][y]==r.Chess[i][j]{
			count++
		}else{
			break
		}
	}
	for i,j:=x+1,y-1;i<=9&&j>=0;i,j=i+1,j-1{
		if r.Chess[x][y]==r.Chess[i][j]{
			count++
		}else{
			break
		}
	}
	if count>=5{
		return true
	}
	return false
}
