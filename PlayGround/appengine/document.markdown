

# Play with Google App Engine

## Get started

Google's documents for starters.

https://developers.google.com/appengine/docs/go/

https://developers.google.com/appengine/docs/go/gettingstarted


## prerequisites

Python 2.7:

    http://www.python.org/ftp/python/

Or Portable Python

    http://www.portablepython.com/wiki/Download


## Questions

1. Question 1: The 'go' runtime is only supported for apps using the High Replication Datastore

The Python 2.7 and Go runtimes are not supported for applications that use the Master/Slave Datastore.

1.1 More about App Engine Storage Options:

    High Replication (default)

        Uses a more highly replicated Datastore that makes use of a system based on the Paxos algorithm to synchronously replicate data across multiple locations simultaneously. Offers the highest level of availability for reads and writes at the cost of eventual consistency for some queries.

    Master/Slave Deprecated!
        Uses a master-slave replication system, which asynchronously replicates data as you write it to another physical datacenter. Since only one datacenter is the master for writing at any given time, this option offers strong consistency for all reads and queries, at the cost of periods of temporary unavailability during datacenter issues or moves.


2. 让Google App Engine支持HTTPS

在app.yaml中设置handler时，设置secure属性:

    - url: /accounts/.*
      script: admin.py
      login: admin
      secure: always

secure的属性可以设成 "always", "optional", or "never" .


3. How to use appcfg.py ?


From: http://www.keakon.net/2009/02/27/appcfg.py%E4%BD%BF%E7%94%A8%E4%BB%8B%E7%BB%8D


大多数时候，我们可能只会用到appcfg.py的上传整个应用程序的功能。

3.1 上传应用程序

这是最常用的命令，格式如下：

    appcfg.py update myapp/

其中，myapp是应用程序的根目录，请改成你自己的目录名。并且，该目录下要有app.yaml这个文件（否则你的应用程序无法访问）。
之后会检查cookie，如果没找到，则要求你输入该应用程序管理员的Google账号和密码。
你也可以用这种格式来指定账号名，但不能显示指定密码：

appcfg.py --email=Albert.Johnson@example.com update myapp/

　
3.2 更新index

由于更新应用程序时，可能会用到新的查询，该查询又需要新的index，但index的购建需要一个过程。
在这段期间里，如果用户进行了这个查询，可能会导致查询失败。你有2种方式避免这一状况发生：

    - 更改app.yaml的Versions段。

    - 先更新index，再上传新应用程序。

        appcfg.py update_indexes myapp/

你可以在后台看到你的index状态。

3.3 删除不用的index

当你从index.yaml中删除一个index时，应用程序并不会自动删除这个index，以方便你可能想要退回到原版本。
然而用不到的index可能会影响到你的程序性能。如果你确定需要删除已经用不到的index，使用下面的命令即可：

appcfg.py vacuum_indexes myapp/

它会自动判断哪些索引已经不在于index.yaml文件里了。


3.4 下载log

在后台查看logging模块的记录非常麻烦，其实你可以将它下载下来，而且还提供了更详细的信息。执行下面的命令即可：

appcfg.py request_logs myapp/ mylogs.txt

缺省情况下，该命令只会下载当天0点（以太平洋标准时间午夜为准）以来，INFO级别及以上的记录（即不包括DEBUG和REQUEST ONLY），并且会覆盖mylogs.txt这个文件。
你可以通过命令行参数，指定时间（天数）、最小记录等级、覆盖还是附加在mylogs.txt文件。下文的命令行参数会有介绍。

3.5 使用HTTP代理

如果需要通过代理来使用appcfg.py，则需要做以下配置：

    - Windows：

        set HTTP_PROXY=http://cache.mycompany.com:3128
        appcfg.py update myapp

    - Mac OS X和Linux：

        export http_proxy="http://cache.mycompany.com:3128"
        appcfg.py update myapp

　
3.6 常用命令行参数

上传和更新应用程序

    appcfg.py [options] update <app-directory>

撤销一次不完整的更新。当更新中断，提示你更新被lock时，就需要执行这个操作。

    appcfg.py [options] rollback <app-directory>

更新index

    appcfg.py [options] update_indexes <app-directory>

删除不用的index

    appcfg.py [options] vacuum_indexes <app-directory>

下载log数据。如果输入的output-file是连接号（-），则输出在命令控制台。

    appcfg.py [options] request_logs <app-directory> <output-file>

    --num_days=...：获取的天数。如果不设置该参数，默认情况下，如果没给--append参数，则取1天；否则则为0（取全部记录）。
    --severity=...：指定最低的记录级别。 4：CRITICAL, 3：ERROR, 2：WARNING, 1：INFO, 0：DEBUG。缺省值为1。
    --append：指定输出为附加方式。当输出文件为-时无效。

输出命令帮助

    appcfg.py help <action>

所有的命令都能用这些参数：

    --quiet：不显示是否成功或失败。
    --verbose：显示当前动作。
    --noisy：显示当前动作的详细信息。当团队合作时很有用。
    --email=...：设置应用程序管理员的Google账号的email地址。如果cookie中没有记录的话，就会使用该账号。
    --server=...：使用的GAE主机名，默认值是appengine.google.com。
    --host=...：使用RPC时的本地主机名。
    --no_cookies：不保存cookie，每次都要求重新输入密码。
    --force：强制删除不使用的index。
    --max_size=...：设置上传的最大文件大小。GAE限制其为1MB（文档还未更新，目前应该是10MB）。
