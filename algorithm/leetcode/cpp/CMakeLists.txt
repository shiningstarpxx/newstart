project("leetcode")

#add_subdirectory(googletest)
#include_directories(${gtest_SOURCE_DIR}/include ${gtest_SOURCE_DIR})

include(FetchContent)
FetchContent_Declare(
        googletest
        # Specify the commit you depend on and update it regularly.
        URL https://github.com/google/googletest/archive/609281088cfefc76f9d0ce82e1ff6c30cc3591e5.zip
)
# For Windows: Prevent overriding the parent project's compiler/linker settings
set(gtest_force_shared_crt ON CACHE BOOL "" FORCE)
FetchContent_MakeAvailable(googletest)

# 'Google_Tests_run' is the target name
# 'test1.cpp tests2.cpp' are source files with tests
add_executable(Google_Tests_run math_260lt.cpp dfs_488lt.cpp dp_629lt.cpp dp_375lt.cpp simulation_520lt.cpp tire_tree_677lt.cpp
        math_319lt.cpp sweep_line_391lt.cpp bit_operation_318lt.cpp dfs_397lt.cpp hashtable_594lt.cpp multi_tree_traverse_1600lt.cpp
        hashtable_423lt.cpp math_458lt.cpp random_519lt_tle.cpp random_380lt.cpp random_519lt.cpp  priority_queue_786lt.cpp dfs_1034lt.cpp
        slide_window_689lt.cpp simulation_794.cpp binary_search_911.cpp greedy_630.cpp topological_sort_851.cpp dfs_flood_fill_419.cpp
        dfs_flood_fill_419.h simulation_graph_997.cpp binary_search_475.cpp string_kmp_686.cpp string_rabin_karp_1044.cpp
        priority_queue_1705.cpp trie_472.cpp hash_1995.cpp hashmap_846.cpp prefix_sum_724.cpp perfix_sum_560.cpp prefix_sum_930.cpp
        prefix_sum_1248.cpp prefix_sum_974.cpp prefix_sum_523.cpp game_dfs_913.cpp dp_route_plan_62.cpp string_71.cpp
        dp_route_plan_63.cpp dp_route_plan_64.cpp simualtion_math_306.cpp bfs_didirection_1036.cpp multi_merge_373.cpp
        multi_merge_378.cpp dp_1220.cpp string_time_539.cpp bfs_1345.cpp graph_743.cpp graph_743_adj_list.cpp sort_969.cc binary_1601.cpp dp_2100.cpp prefix_sum_2100.cpp prefix_sum_2055.cpp difference_array_798.cpp difference_array_1109.cpp dfs_2049.cpp data_structure_432.cpp)
target_link_libraries(Google_Tests_run gtest gtest_main)