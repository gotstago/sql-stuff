package main

import (
	"fmt"
	"regexp"
)

func main() {
	//var re = regexp.MustCompile(`(?P<time>\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2},\d{3}) (?P<level>\S*) (?P<processor>\S*) (?P<serviceLayer>\S*):(?P<lineNumber>\d*) - (?P<mess>.*)`)
	// var pre = regexp.MustCompile(`(?P<time>\d{4} \S{3} \d{2} \d{2}:\d{2}:\d{2}) .* (?P<level>\S*) |.* (?P<class>\S*) .* (?P<unknown>\S*) .* (?P<mess>.*)`)
	var pre = regexp.MustCompile(`(?sm)(\d{4} \S{3} \d{2} \d{2}:\d{2}:\d{2}\s\S\sERROR.*?)(?:\d{4} \S{3})`)
	//var pre = regexp.MustCompile(`(\d{4} \S{3} \d{2}) (\d{2}:\d{2}:\d{2}) \[(.*)\] ([^ ]*) +([^ ]*) (.*)$`)
	// var pre = regexp.MustCompile(`(?sm)^\d{4}\s+\Sd{3}\s+\d{2}\s+\S+\sERROR.*?\n\d{4}\s+\Sd{3}\s+\d{2}`)
	var pstr = `2017 Dec 14 15:22:32 | ERROR | com.datacert.core.cognos.service.impl.CognosAPIDeployment | admin | zZJhw9n6BNIt | An error occurred while querying CM object:
	; nested exception is: 
		java.net.ConnectException: Connection refused: connect
	2017 Dec 14 15:22:36 | INFO  | com.datacert.core.common.dao.impl.OSIVInterceptor | admin | zZJhw9n6BNIt | Session token=8lM0DRwtek0I, parameter token=null
	2017 Dec 14 15:22:41 | INFO  | com.datacert.core.common.dao.impl.OSIVInterceptor | admin | zZJhw9n6BNIt | Session token=8lM0DRwtek0I, parameter token=null
	2017 Dec 14 15:22:41 | INFO  | com.datacert.core.common.dao.impl.OSIVInterceptor | admin |  | Session token=8lM0DRwtek0I, parameter token=null
	2017 Dec 14 15:22:41 | INFO  | com.datacert.core.common.dao.impl.OSIVInterceptor | admin | zZJhw9n6BNIt | Session token=8lM0DRwtek0I, parameter token=null
	2017 Dec 14 15:22:41 | INFO  | com.datacert.core.common.dao.impl.OSIVInterceptor | admin |  | Session token=8lM0DRwtek0I, parameter token=null
	2017 Dec 14 15:22:46 | INFO  | com.datacert.core.common.dao.impl.OSIVInterceptor | admin | zZJhw9n6BNIt | Session token=8lM0DRwtek0I, parameter token=null
	2017 Dec 14 15:15:19 | INFO  | com.datacert.core.spring.CustomXmlWebApplicationContext |  |  | Refreshing Root WebApplicationContext: startup date [Thu Dec 14 15:15:19 UTC 2017]; root of context hierarchy
	2017 Dec 14 15:15:22 | ERROR  | com.datacert.core.spring.CustomDefaultListableBeanFactory |  |  | Overriding bean definition for bean 'groovyEntityClassProvider': replacing [Generic bean: class [com.datacert.core.entity.model.plugins.impl.GroovyEntityClassProvider]; scope=singleton; abstract=false; lazyInit=false; autowireMode=0; dependencyCheck=0; autowireCandidate=true; primary=false; factoryBeanName=null; factoryMethodName=null; initMethodName=null; destroyMethodName=null; defined in URL [jar:file:/C:/Datacert/Passport/tomcat/webapps/Passport/WEB-INF/lib/framework-java-2.5.0.187.283673.jar!/com/datacert/core/entity/model/plugins/impl/GroovyEntityClassProvider.class]] with [Generic bean: class [com.datacert.core.entity.model.plugins.impl.GroovyEntityClassProvider]; scope=; abstract=false; lazyInit=false; autowireMode=0; dependencyCheck=0; autowireCandidate=true; primary=false; factoryBeanName=null; factoryMethodName=null; initMethodName=null; destroyMethodName=null; defined in class path resource [core-applicationContext.xml]]
	2017 Dec 14 15:15:23 | INFO  | com.datacert.core.spring.CustomDefaultListableBeanFactory |  |  | Overriding bean definition for bean 'propertyConfigurer': replacing [Generic bean: class [org.springframework.beans.factory.config.PropertyPlaceholderConfigurer]; scope=; abstract=false; lazyInit=false; autowireMode=0; dependencyCheck=0; autowireCandidate=true; primary=false; factoryBeanName=null; factoryMethodName=null; initMethodName=null; destroyMethodName=null; defined in ServletContext resource [/WEB-INF/config/core/applicationContext.xml]] with [Generic bean: class [com.datacert.core.util.AimsDecryptingPropertyPlaceholderConfigurer]; scope=; abstract=false; lazyInit=false; autowireMode=0; dependencyCheck=0; autowireCandidate=true; primary=false; factoryBeanName=null; factoryMethodName=null; initMethodName=null; destroyMethodName=null; defined in ServletContext resource [/WEB-INF/applicationContext.xml]]
	2017 Dec 14 15:15:25 | INFO  | com.datacert.core.util.AimsDecryptingPropertyPlaceholderConfigurer |  |  | Loading properties file from class path resource [jdbc.properties]
	2017 Dec 14 15:15:25 | INFO  | com.datacert.core.util.AimsDecryptingPropertyPlaceholderConfigurer |  |  | Loading properties file from class path resource [email.properties]`

	// 	var str = `2014-08-29 14:53:58,948 DEBUG ajp-bio-172.16.2.157-8009-exec-111 PostgresIdsDAOImpl#getDatabaseConnection:822 - Attempting to JDBC connect to the database
	// 2014-08-29 14:53:58,949 DEBUG ajp-bio-172.16.2.157-8009-exec-111 PostgresIdsDAOImpl#getDatabaseConnection:833 - JDBC Driver is registered successfully
	// 2014-08-29 14:53:58,954 DEBUG ajp-bio-172.16.2.157-8009-exec-111 PostgresIdsDAOImpl#select:389 - Database connection successfully closed.
	// 2014-08-29 14:53:58,954 DEBUG ajp-bio-172.16.2.157-8009-exec-111 ContentManagementServiceImpl#getNiseImages:915 - Postgres returned 0 pix result(s).
	// 2014-08-29 14:53:58,955 DEBUG ajp-bio-172.16.2.157-8009-exec-111 ContentManagementServiceImpl#getNiseImages:923 - Could not find any images after image ID 94611681 limit by 500
	// 2014-08-29 14:53:58,955 DEBUG ajp-bio-172.16.2.157-8009-exec-111 IdsServiceImpl#getNiseImages:2687 - Failed to find any NISE URLs
	// 2014-08-29 14:53:58,955 DEBUG ajp-bio-172.16.2.157-8009-exec-111 EventsServiceImpl#addUserActivity:82 - There are 1 user activities to be inserted.
	// 2014-08-29 14:53:58,955 DEBUG ajp-bio-172.16.2.157-8009-exec-111 RestResource#getNewImages:267 - No new images found during this period
	// 2014-08-29 14:54:00,002 DEBUG pool-2-thread-1 SessionServiceImpl#fetchBannedIps:201 - Banned Ips task is running now
	// 2014-08-29 14:54:00,002 DEBUG pool-2-thread-1 PostgresIdsDAOImpl#select:298 - Finding any number of entries in the DB table
	// 2014-08-29 14:54:00,002 DEBUG pool-2-thread-1 PostgresIdsDAOImpl#select:351 - Search query: SELECT  *  FROM ONLY bannedips WHERE id is not null;
	// 2014-08-29 14:54:00,002 DEBUG pool-2-thread-1 PostgresIdsDAOImpl#getDatabaseConnection:822 - Attempting to JDBC connect to the database
	// 2014-08-29 14:54:00,002 DEBUG pool-2-thread-1 PostgresIdsDAOImpl#getDatabaseConnection:833 - JDBC Driver is registered successfully
	// 2014-08-29 14:54:00,360 DEBUG pool-2-thread-1 PostgresIdsDAOImpl#select:389 - Database connection successfully closed.
	// 2014-08-29 14:54:00,360 DEBUG pool-2-thread-1 SessionServiceImpl#fetchBannedIps:208 - Postgres returned 31164 banned ips
	// 2014-08-29 14:54:00,418 DEBUG pool-2-thread-1 SessionServiceImpl#fetchBannedIps:222 - There are 31164 banned ips
	// 2014-08-29 14:54:20,002 DEBUG pool-2-thread-1 UserServiceImpl#fetchLoginConfig:234 - Fetch the login configuration for banning IPs
	// 2014-08-29 14:54:20,002 DEBUG pool-2-thread-1 PostgresIdsDAOImpl#select:298 - Finding any number of entries in the DB table
	// 2014-08-29 14:54:20,002 DEBUG pool-2-thread-1 PostgresIdsDAOImpl#select:351 - Search query: SELECT  *  FROM ONLY sitesloginparameters WHERE usernamefail is not null;
	// 2014-08-29 14:54:20,002 DEBUG pool-2-thread-1 PostgresIdsDAOImpl#getDatabaseConnection:822 - Attempting to JDBC connect to the database
	// 2014-08-29 14:54:20,002 DEBUG pool-2-thread-1 PostgresIdsDAOImpl#getDatabaseConnection:833 - JDBC Driver is registered successfully
	// 2014-08-29 14:54:20,008 DEBUG pool-2-thread-1 PostgresIdsDAOImpl#select:389 - Database connection successfully closed.
	// 2014-08-29 14:54:20,008 DEBUG pool-2-thread-1 UserServiceImpl#fetchLoginConfig:241 - Postgres returned 1 login config setting
	// 2014-08-29 14:54:20,008 DEBUG pool-2-thread-1 UserServiceImpl#fetchLoginConfig:255 - LoginConfig Parameter is updated true
	// 2014-08-29 14:54:30,002 DEBUG pool-2-thread-1 EventsServiceImpl#processUserActivities:49 - User activities task is running now
	// 2014-08-29 14:54:30,002 DEBUG pool-2-thread-1 EventsServiceImpl#processUserActivities:56 - There are 1 user activities to be inserted
	// 2014-08-29 14:54:30,002 DEBUG pool-2-thread-1 PostgresIdsDAOImpl#select:82 - Finding any number of entries in the DB table
	// 2014-08-29 14:54:30,002 DEBUG pool-2-thread-1 PostgresIdsDAOImpl#select:131 - Search query: SELECT id,text FROM user_agents WHERE text='NISE-IDS solr-b04 14000';
	// 2014-08-29 14:54:30,002 DEBUG pool-2-thread-1 PostgresIdsDAOImpl#getDatabaseConnection:822 - Attempting to JDBC connect to the database
	// 2014-08-29 14:54:30,002 DEBUG pool-2-thread-1 PostgresIdsDAOImpl#getDatabaseConnection:833 - JDBC Driver is registered successfully
	// 2014-08-29 14:54:30,011 DEBUG pool-2-thread-1 PostgresIdsDAOImpl#select:175 - Database connection successfully closed.
	// 2014-08-29 14:54:30,011 DEBUG pool-2-thread-1 EventsServiceImpl#getUserAgent:479 - Postgres returned 2 result(s).
	// 2014-08-29 14:54:30,011 DEBUG pool-2-thread-1 PostgresIdsDAOImpl#insert:403 - Inserting new entry into Postgres DB table
	// 2014-08-29 14:54:30,011 DEBUG pool-2-thread-1 PostgresIdsDAOImpl#getDatabaseConnection:822 - Attempting to JDBC connect to the database
	// 2014-08-29 14:54:30,011 DEBUG pool-2-thread-1 PostgresIdsDAOImpl#getDatabaseConnection:833 - JDBC Driver is registered successfully
	// 2014-08-29 14:54:30,016 DEBUG pool-2-thread-1 PostgresIdsDAOImpl#insert:449 - Database connection successfully closed.
	// 2014-08-29 14:54:30,016 DEBUG pool-2-thread-1 EventsServiceImpl#insertUserActivity:454 - 1 user activity inserted
	// 2014-08-29 14:54:30,016 DEBUG pool-2-thread-1 EventsServiceImpl#processUserActivities:70 - At the end of thread there are 0 user activities in the list
	// 2014-08-29 14:55:00,001 DEBUG pool-2-thread-1 SessionServiceImpl#fetchBannedIps:201 - Banned Ips task is running now
	// 2014-08-29 14:55:00,001 DEBUG pool-2-thread-1 PostgresIdsDAOImpl#select:298 - Finding any number of entries in the DB table
	// 2014-08-29 14:55:00,002 DEBUG pool-2-thread-1 PostgresIdsDAOImpl#select:351 - Search query: SELECT  *  FROM ONLY bannedips WHERE id is not null;
	// 2014-08-29 14:55:00,002 DEBUG pool-2-thread-1 PostgresIdsDAOImpl#getDatabaseConnection:822 - Attempting to JDBC connect to the database`
	// n1 := pre.SubexpNames()
	// r2 := pre.FindAllStringSubmatch(pstr, -1)[0]

	// fmt.Printf("The names are  : %v\n", n1)
	// fmt.Printf("The matches are: %v\n", r2)
	for i, match := range pre.FindAllStringSubmatch(pstr, -1) {
		fmt.Println(match[1], "found at index", i)
	}
}
