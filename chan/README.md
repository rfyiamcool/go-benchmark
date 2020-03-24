## benchmark result

aliyun 8core 2.5hz

```
direct  count: 20000000, chan buf: 0, sender: 10, recevier: 10, time cost: 26.8407462s, qps: 745135.77
dispatch count: 20000000, chan buf: 0, sender: 10, recevier: 10, time cost: 5.976753845s, qps: 3346298.20
direct  count: 20000000, chan buf: 50, sender: 10, recevier: 10, time cost: 14.091552977s, qps: 1419290.01
dispatch count: 20000000, chan buf: 50, sender: 10, recevier: 10, time cost: 1.781377337s, qps: 11227268.34
direct  count: 20000000, chan buf: 100, sender: 10, recevier: 10, time cost: 13.05101978s, qps: 1532447.34
dispatch count: 20000000, chan buf: 100, sender: 10, recevier: 10, time cost: 1.457061197s, qps: 13726261.05
direct  count: 20000000, chan buf: 1000, sender: 10, recevier: 10, time cost: 11.760349875s, qps: 1700629.73
dispatch count: 20000000, chan buf: 1000, sender: 10, recevier: 10, time cost: 826.241882ms, qps: 24205995.63

direct  count: 20000000, chan buf: 0, sender: 50, recevier: 50, time cost: 27.134798244s, qps: 737060.95
dispatch count: 20000000, chan buf: 0, sender: 50, recevier: 50, time cost: 5.753107347s, qps: 3476382.38
direct  count: 20000000, chan buf: 100, sender: 50, recevier: 50, time cost: 13.928744062s, qps: 1435879.66
dispatch count: 20000000, chan buf: 100, sender: 50, recevier: 50, time cost: 1.887928385s, qps: 10593625.43
direct  count: 20000000, chan buf: 1000, sender: 50, recevier: 50, time cost: 11.869635307s, qps: 1684971.77
dispatch count: 20000000, chan buf: 1000, sender: 50, recevier: 50, time cost: 889.46773ms, qps: 22485366.44

direct  count: 20000000, chan buf: 0, sender: 100, recevier: 100, time cost: 27.125478349s, qps: 737314.19
dispatch count: 20000000, chan buf: 0, sender: 100, recevier: 100, time cost: 5.522589581s, qps: 3621489.61
direct  count: 20000000, chan buf: 100, sender: 100, recevier: 100, time cost: 13.483081082s, qps: 1483340.49
dispatch count: 20000000, chan buf: 100, sender: 100, recevier: 100, time cost: 2.007464582s, qps: 9962817.04
direct  count: 20000000, chan buf: 1000, sender: 100, recevier: 100, time cost: 11.939463813s, qps: 1675117.11
dispatch count: 20000000, chan buf: 1000, sender: 100, recevier: 100, time cost: 1.223054583s, qps: 16352503.98

direct  count: 20000000, chan buf: 0, sender: 200, recevier: 200, time cost: 26.721733643s, qps: 748454.44
dispatch count: 20000000, chan buf: 0, sender: 200, recevier: 200, time cost: 5.683547095s, qps: 3518929.40
direct  count: 20000000, chan buf: 100, sender: 200, recevier: 200, time cost: 14.318196445s, qps: 1396823.98
dispatch count: 20000000, chan buf: 100, sender: 200, recevier: 200, time cost: 2.1512151s, qps: 9297072.39
direct  count: 20000000, chan buf: 1000, sender: 200, recevier: 200, time cost: 11.969512866s, qps: 1670911.81
dispatch count: 20000000, chan buf: 1000, sender: 200, recevier: 200, time cost: 1.329972017s, qps: 15037913.61

direct  count: 20000000, chan buf: 0, sender: 300, recevier: 300, time cost: 26.94974238s, qps: 742122.13
dispatch count: 20000000, chan buf: 0, sender: 300, recevier: 300, time cost: 5.707497569s, qps: 3504162.77
direct  count: 20000000, chan buf: 100, sender: 300, recevier: 300, time cost: 13.920845708s, qps: 1436694.35
dispatch count: 20000000, chan buf: 100, sender: 300, recevier: 300, time cost: 1.499986318s, qps: 13333457.13
direct  count: 20000000, chan buf: 1000, sender: 300, recevier: 300, time cost: 12.084544618s, qps: 1655006.54
dispatch count: 20000000, chan buf: 1000, sender: 300, recevier: 300, time cost: 1.083441217s, qps: 18459705.93
```
