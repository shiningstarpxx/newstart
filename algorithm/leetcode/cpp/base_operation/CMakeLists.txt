project("base_operation")

include(FetchContent)
set(CMAKE_CXX_STANDARD 14)

# add_executable(set_operation set_operation.cpp)
# add_executable(priority_queue_operation ./priority_queue_operation/priority_queue_operation.cpp ../greedy_priority_queue_and_set_1606.cpp ../kadane_918.cpp ../slide_window_3.cpp tree_array_fenwich_tree/fenwick_tree.cpp)

# 如果你只需要根目录下的 test 文件夹的所有 .cpp 文件
#file (GLOB files test/*.cpp)
# 如果你只有两层目录的话
#file (GLOB files *.cpp */*.cpp)
# 同理，三层的话
#file (GLOB files *.cpp */*.cpp */*/*.cpp)

# 官方提供了一种递归的方法
# 这样在运行框会多一个 CMakeCXXCompilerId，不过无伤大雅
# 遍历项目根目录下所有的 .cpp 文件
file (GLOB_RECURSE files *.cpp)
foreach (file ${files})
    string(REGEX REPLACE ".+/(.+)\\..*" "\\1" exe ${file})
    add_executable (${exe} ${file} ../array_tree.cpp ../dp_673.cpp ../array_tree_1310.cpp ../dp_139.cpp ../dp_42.cpp ../dp_413.cpp ../bfs_310.cpp presum.cpp ../dp_route_plan_221.cpp ../dp_distance_72.cpp ../dp_combination_sum_416.cpp ../dp_902.cpp ../dp_string_131.cpp knapsack/knapsack.cpp ../dfs_417.cpp ../slide_window_713.cc ../dp_1728.cpp ../dp_300.cpp ../dp_354.cpp ../dp_918.cpp ../dp_matrix_sum_17.24.cpp ../dp_matrix_sum_363.cpp ../toplogic_sort_offer:_114.cpp ../dp_473.cc ../tree_450.cpp cmake-build-debug/segment_tree.cpp ../setment_tree_307.cpp ../dp_1deminsion_withk_813.cpp)
    message (\ \ \ \ --\ src/${exe}.cpp\ will\ be\ compiled\ to\ bin/${exe})
endforeach ()