package array

/* 数组 arrayA 记录了各个生物群体数量数据，其中 arrayA[i] 表示第 i 个生物群体的数量。
请返回一个数组 arrayB，该数组为基于数组 arrayA 中的数据计算得出的结果，其中 arrayB[i] 表示将第 i 个生物群体的数量从总体中排除后的其他数量的乘积

输入：arrayA = [2, 4, 6, 8, 10]
输出：[1920, 960, 640, 480, 384]
 */


 // StatisticalResult 一维转二维（很多都适用) 
 //  @param arrayA 
 //  @return []int 
 func StatisticalResult(arrayA []int) []int {
    if len(arrayA)==0{
        return []int{}
    }

	//先迭代求下半角，再迭代求上半角
    arrayB:=make([]int,len(arrayA))
    arrayB[0]=1
    for i:=1;i<len(arrayA);i++{
        arrayB[i]=arrayB[i-1]*arrayA[i-1]
    }

    temp:=1
    for i:=len(arrayA)-2;i>=0;i--{
        temp=temp*arrayA[i+1]
        arrayB[i]=arrayB[i]*temp
    }

    return arrayB
}