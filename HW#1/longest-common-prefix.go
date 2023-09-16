func longestCommonPrefix(s []string) string {
    if len(s) == 0 {
        return ""
    }

    ans := s[0]

    for i := 1; i < len(s); i++ {
        t := ""
        for j := 0; j < len(s[i]); j++ {
            if j >= len(ans) || ans[j] != s[i][j] {
                break
            }
            t += string(s[i][j])
        }
        ans = t
    }
    return ans
}