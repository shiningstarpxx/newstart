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
add_executable(Google_Tests_run math_260lt.cpp dfs_488lt.cpp dp_629lt.cpp dp_375lt.cpp simulation_520lt.cpp tire_tree_677lt.cpp math_319lt.cpp sweep_line_391lt.cpp bit_operation_318lt.cpp dfs_397lt.cpp hashtable_594lt.cpp multi_tree_traverse_1600lt.cpp hashtable_423lt.cpp math_458lt.cpp random_519lt_tle.cpp random_380lt.cpp random_519lt.cpp)
target_link_libraries(Google_Tests_run gtest gtest_main)