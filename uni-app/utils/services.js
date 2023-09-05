import config from '../config';

// 配置 APPServer 服务域名
export const ServicesOrigin = config.appServer;

// 配置api接口路径前缀
export const ApiPrefixPath = '/api/v1/live/';

const UserInfoStorageKey = 'user_info';

class Services {
	constructor() {
		this.userId = '';
		this.userName = '';
		this.authToken = '';
		this.expire = '';
		
		// 从缓存中读用户数据
		try{
			const infoStr = uni.getStorageSync(UserInfoStorageKey);
			if (infoStr) {
				const info = JSON.parse(infoStr);
				this.userId = info.userId;
				this.userName = info.userName;
				this.authToken = info.authToken;
				this.expire = info.expire;
			}
		}catch(e){
			//TODO handle the exception
		}
	}
	
	getUserInfo() {
		return {
			userId: this.userId,
			userName: this.userName,
		};
	}
	
	request(api, data, header) {
		const url = ServicesOrigin + ApiPrefixPath + api;
		// console.log('url-->', url);
		return new Promise((resolve, reject) => {
			uni.request({
				url,
				method: 'POST',
				data,
				header,
				success(res) {
					if (res.statusCode === 200 && res.data) {
						resolve(res.data);
						return;
					}
					if (res.statusCode === 401) {
						// 处理登录态失效等问题
					}
					reject(res.data);
				},
				fail(err) {
					reject(err);
				}
			})
		});
	}
	
	login(userId, username) {
		// 实际场景请勿明文传密码
		return this.request(
			'login',
			{
				username,
				password: username,
			},
		).then((res) => {
			this.userId = userId;
			this.userName = username;
			this.authToken = res.token;
			this.expire = res.expire;
			
			uni.setStorage({
				key: UserInfoStorageKey,
				data: JSON.stringify({
					userId,
					userName: username,
					authToken: res.token,
					expire: res.expire,
				}),
			});
		});
	}
	
	getRoomList(pageNum, pageSize) {
		return this.request(
			'list',
			{
				user_id: this.userId,
				page_num: pageNum,
				page_size: pageSize,
			},
			{
				Authorization: `Bearer ${this.authToken}`,
			}
		);
	}
	
	getRoomDetail(roomId) {
		return this.request(
			'get',
			{
				user_id: this.userId,
				id: roomId,
			},
			{
				Authorization: `Bearer ${this.authToken}`,
			}
		);
	}
	
	getToken() {
		return this.request(
			'token',
			{
				user_id: this.userId,
				device_type: 'web',
				device_id: 'isuniapp',
			},
			{
				Authorization: `Bearer ${this.authToken}`,
			}
		);
	}
}

export default new Services();
