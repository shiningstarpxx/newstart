/**
 * @Author: pei, xingxin
 * @mail: xingxinpei@gmail.com
 * @Description:
 * @File: decode
 * @Version: 1.0.0
 * @Data: 2022/3/15 8:49 PM
 */
package util

// inline uint32_t DecodeFixed32(const char* ptr) {
//  const uint8_t* const buffer = reinterpret_cast<const uint8_t*>(ptr);
//
//  // Recent clang and gcc optimize this to a single mov / ldr instruction.
//  return (static_cast<uint32_t>(buffer[0])) |
//         (static_cast<uint32_t>(buffer[1]) << 8) |
//         (static_cast<uint32_t>(buffer[2]) << 16) |
//         (static_cast<uint32_t>(buffer[3]) << 24);
//}

func DecodeFixed32(ptr string) uint32 {
	if len(ptr) != 4 {
		panic("fixer32 decode must 4")
	}
	return uint32(ptr[0]) | (uint32(ptr[1]) << 8) | (uint32(ptr[2]) << 16) | (uint32(ptr[3]) << 24)
}