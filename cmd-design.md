1. 列出命令说明
/*

Post: 列出所有命令以及其说明

*/
agenda help


2. 列出某个命令的描述
/*

Pre: 提供某个命令的名称

Post: 列出该命令的所有描述

*/
agenda help xxx


3. 用户注册
/*

Pre: 提供唯一的用户名、密码、邮箱、电话用来注册用户

Post: 如果注册成功，显示 "Succeed to Create User xxx!";　如果注册失败，只有1种情况：　１）用户已经存在，显示 "Fail to Create User! User already exists!"

*/
agenda register -u/-username Username -p/--password Password -e/--email Email -t/--telephone Telephone


4. 用户登录
/*

Pre: 提供用户名以及密码进行登录

Post:　如果用户名和密码正确，登陆成功，显示 "Succeed to Log in!"; 如果登录失败: 分3种情况：　１)用户不存在，显示 "Fail to Login in! User doesn't exist!"  2)用户名与密码不匹配，显示　"Fail to Login in! The password doesn't match the username!"  3)用户已经登录，显示 "Fail to Login in! User is already log in!"

*/
agenda login -u/--username Username -p/--password Password


5. 用户登出
/*

Pre: 用户已经登录；用户登出后只能进行用户注册和用户等录

Post: 登出成功，显示 "Succeed to Log out!";  登出失败，有2种情况   1）用户不存在，显示 "Fail to Log out! User hasn't Register yet!"   2)用户没有登录，显示 "Fail to Log out! User hasn't Loged in yet!"

*/
agenda logout --username username


6. 用户查询
/*

Pre: 用户已经注册；提供注册用户的用户名、邮箱或者电话，只需其中一个就行，但如果三个都提供了，那么将三个作为独立的查询条件，然后将三个查询的结果合并返回；或者什么都不提供（查看所有注册用户）

Post: 如果不带有option选项，就显示所有注册用户； 如果带有option选项，如果该用户已经注册，就显示该用户的信息；如果该用户没有注册，那么就显示 "Fail to search User xxx! User xxx hasn't register yet!"

*/
agenda searchUser -u/--username username -i/--information [Username/Email/Telephone]


7. 用户删除
/*

Pre: 用户已经登录；只能删除自己的账号

Post: 删除成功：显示　"Succeed to Delete User xxx!" 同时以该用户为发起者的会议将被删除；以该用户为参与者的会议将从参与者列表中移除该用户，若因此造成会议参与者人数变为０，则该会议也会被删除；　　删除失败：  有2种情况  １）用户不存在，显示 "Fail to Delete User xxx! User xxx hasn't Register yet!"　2)用户还没登录，显示 "Fail to Delete User xxx! User xxx hasn't Log in yet!"

*/
agenda deleteUser --username username


8. 创建会议
/*

Pre: 用户已经登录；不能包括未注册的用户；提供会议的主题、参与者、起始时间、结束时间这四个信息；

Post: 创建成功：显示 "User xxx Succeed to Create Meeting xxxx!" 创建失败：有以下4种情况：　１)用户还没注册，显示 "User xxx Fail to Create Meeting! User xxx hasn't register yet!"  2)用户没有登录，显示 "User xxx Fail to Create Meeting! User xxx hasn't Log in yet!"  3)会议已经存在，显示　"User xxx Fail to Create Meeting! Meeting has been Created!"  4)参与者并没有全部都是注册的，显示 User xxx Fail to Create Meeting! [xxx, xxx] hasn't register yet!  5)无论是作为发起者还是参与者，如果与现有会议时间重叠，创建会议就会失败，显示 "User xxx Fail to Create Meeting! [xxx, xxx, ...] can't not participate the meeting during this period!"

*/
agenda cm --username username --title Title --participator [xxx,xxx, ...] --startTime StartTime --endTime EndTime


9. 增加会议参与者
/*

Pre: 用户已经登录；自己发起的会议

Post: 增加成功：显示 "Succeed to Add Participator [xxx, ...] to the Meeting xxxx!"  增加失败: 有6种情况：　１)用户不存在，显示 "Fail to Add Participator [xxx, ...] to the Meeting! User hasn't Register yet!"  2)用户没有登录，显示 "Fail to Add Participator [xxx, ...] to the Meeting! User hasn't Log in yet!"  3)会议不存在，显示 "Fail to Add Participator [xxx, ...] to the Meeting! Meeting hasn't been Created  yet!"  4)用户并没有创建该会议，也就是用户没有权利去增删参与者，显示 "Fail to Add Participator [xxx, ...] to the Meeting! User has no right to add participator!"   5)参与者还没有进行注册，显示 "Fail to Add Participator [xxx, ...] to the Meeting! Participator [xxx, ...] hasn't register yet!"  6)该会议和参与者其他会议的时间发生重叠，显示 "Fail to Add Participator [xxx, ...] to the Meeting! Participator [xxx, ...] can't not participate the meeting during this period!"

*/
agenda addpar --username username --title Title --participator [xxx, ...]


10. 删除会议参与者
/*

Pre: 用户已经登录；自己发起的会议

Post: 删除成功：显示 "Succeed to Delete Participator [xxx, ...] from the Meeting xxxx!" 同时如果该会议的参与者变为０，就删除该会议　　 删除失败：　有5种情况：　１)用户不存在，显示 "Fail to Delete Participator [xxx, ...] from the Meeting! User hasn't Register yet!"   2)用户没有登录，显示 "Fail to Delete Participator [xxx, ...] from the Meeting! User hasn't Log in yet!"  3)会议不存在，显示 "Fail to Delete Participator [xxx, ...] from the Meeting! Meeting hasn't been Created yet!"  4)用户并没有创建该会议，也就是用户没有权利去增删参与者，显示 "Fail to Delete Participator [xxx, ...] from the Meeting! User has no right to delete participator!"  5)在自己创建的会议中删除自己（此时会议人数大于1），这里默认是不能退出的，当人数大与1时，只能执行取消操作；但等于1时可以删除，显示 "Fail to Delete Participator [xxx, ...] from the Meeting! User is the organizer of this meeting and the number of the meeting is greater than 1!"

*/
agenda deletepar --username username --title Title --participator [xxx, ...]


11. 查询会议
/*

Pre: 用户已经登录；提供起始时间和结束时间

Post:  查询成功: 显示 "Succeed to Search the Meeting during this period!" 同时将返回列表中的会议列出来　　查询失败: 有３种情况：　１)用户并不存在，显示 "Fail to Search the Meeting! User hasn't Register yet!"   2)用户没有登录，显示 "Fail to Search the Meeting! User hasn't Log in yet!"  3)在时间段内没有找到任何会议，显示 "Fail to Search the Meeting! There is no meeting during this period!"

*/
agenda searchMeeting --username username --startTime StartTime --endTime EndTime


12. 取消会议
/*

Pre: 用户已经登录；用户自己发起的会议

Post: 取消成功: 显示 "Succeed to Cancel the Meeting xxxx!"  取消失败: 有4种情况：　１)用户并不存在，显示 "Fail to Cancel the Meeting! User hasn't Register yet!"  2)用户没有登录，显示 "Fail to Cancel the Meeting! User hasn't Log in yet!"  3)会议还没有创建，显示 "Fail to Cancel the Meeting! Meeting hasn't been created yet!"  4)会议并不是用户创建的，显示 "Fail to Cancel the Meeting! User has no right to cancel this meeting!"

*/
agenda cancel --username --title Title


13. 退出会议
/*

Pre: 用户已经登录；自己参与其中的会议

Post: 退出成功: 显示 "Succeed to Quit the Meeting xxxx!" 如果该会议的参与者人数变为０，则删除该会议　 退出失败: 有5种情况：　１)用户并不存在，显示 "Fail to Quit the Meeting! User hasn't register yet!"  2)用户没有登录，显示 "Fail to Quit the Meeting! User hasn't Log in yet!"  3)会议并不存在，显示 "Fail to Quit the Meeting! Meeting hasn't been created yet!"   4)用户并不在参与者列表中，显示 "Fail to Quit the Meeting! User hasn't participate this meeting yet!"  5)退出自己创建的会议（此时会议人数大于1），这里默认是不能退出的，当人数大与1时，只能执行取消操作；但人数等于1时可以，显示 "Fail to Quit the Meeting! User is the organizer of this meeting and the number of the meeting is greater than 1!"
 
*/
agenda quit --username username --title Title


14. 清空会议
/*

Pre: 用户已经登录；自己发起的所有会议

Post: 清空成功: 显示 "Succeed to Clear all Meetings created by xxx!"   清空失败: １)用户并不存在，显示 "Fail to Clear all Meetings! User hasn't register yet!"  2）用户没有登录，显示 "Fail to Clear all Meetings! User hasn't Log in yet!"

*/
agenda clear --username username





