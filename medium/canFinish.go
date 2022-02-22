package medium

/**
207. 课程表
你这个学期必须选修 numCourses 门课程，记为 0 到 numCourses - 1 。

在选修某些课程之前需要一些先修课程。 先修课程按数组 prerequisites 给出，其中 prerequisites[i] = [ai, bi] ，表示如果要学习课程 ai 则 必须 先学习课程  bi 。

例如，先修课程对 [0, 1] 表示：想要学习课程 0 ，你需要先完成课程 1 。
请你判断是否可能完成所有课程的学习？如果可以，返回 true ；否则，返回 false 。

 

示例 1：

输入：numCourses = 2, prerequisites = [[1,0]]
输出：true
解释：总共有 2 门课程。学习课程 1 之前，你需要完成课程 0 。这是可能的。
示例 2：

输入：numCourses = 2, prerequisites = [[1,0],[0,1]]
输出：false
解释：总共有 2 门课程。学习课程 1 之前，你需要先完成​课程 0 ；并且学习课程 0 之前，你还应先完成课程 1 。这是不可能的。
 

提示：

1 <= numCourses <= 105
0 <= prerequisites.length <= 5000
prerequisites[i].length == 2
0 <= ai, bi < numCourses
prerequisites[i] 中的所有课程对 互不相同

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/course-schedule
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func canFinish(numCourses int, prerequisites [][]int) bool {
	if prerequisites == nil || len(prerequisites) < 2 {
		return true
	}
	//1、计算入度表，并存储节点连接映射表
	inCount := make([]int, numCourses)          //入度表
	nodeList := make(map[int][]int, numCourses) //节点连接映射表
	for i := 0; i < len(prerequisites); i++ {
		inCount[prerequisites[i][0]]++
		lists, ok := nodeList[prerequisites[i][1]]
		if ok {
			lists = append(lists, prerequisites[i][0])
		} else {
			lists = []int{prerequisites[i][0]}
		}
		nodeList[prerequisites[i][1]] = lists
	}
	//2、将所有入度为0的节点 进入队列
	queue := make([]int, 0)
	for i := 0; i < len(inCount); i++ {
		if inCount[i] == 0 {
			queue = append(queue, i)
		}
	}
	//3、DFS，进行拓扑排序
	for len(queue) > 0 {
		//头元素 出队
		pre := queue[0]
		queue = queue[1:]
		numCourses--
		nodes, ok := nodeList[pre]
		if !ok {
			continue
		}
		for _, node := range nodes { //入度为0的节点 指向的节点遍历
			inCount[node] -= 1      //由于父节点进行计算且出队，因此入度-1
			if inCount[node] == 0 { //入度为0，入队
				queue = append(queue, node)
			}
		}
	}
	return numCourses == 0
}
