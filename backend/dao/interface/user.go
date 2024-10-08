/**
 ******************************************************************************
 * @file           : user.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/8/14
 ******************************************************************************
 */

package _interface

type UserDaoInterface interface {
	// SELECT * From user as u1
	// WHERE ui.id IN (	SELECT fo
	// )
	FindFriendList(id uint) ([]uint, error)
}
