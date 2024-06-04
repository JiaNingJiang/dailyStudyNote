package com.buptjjn.test;

import com.buptjjn.ioc_03.HappyComponent;
import com.buptjjn.ioc_05.JavaBean;
import com.buptjjn.ioc_05.JavaBeanFactory;
import org.junit.Test;
import org.springframework.context.ApplicationContext;
import org.springframework.context.support.ClassPathXmlApplicationContext;

public class SpringIocTest {

    /**
     * 创建 ioc 容器，并且读取配置文件
     * 可以指定一个或者多个配置文件
     */
    public ApplicationContext createIoc() {
        // 方式一：创建 ioc 容器的同时读取配置文件
        ApplicationContext context = new ClassPathXmlApplicationContext("spring-ioc-xml-03.xml");
        // 方式二：1. 创建 ioc 容器  2. 读取配置文件  3. 刷新 ioc 容器
        ClassPathXmlApplicationContext context1 = new ClassPathXmlApplicationContext();
        context1.setConfigLocation("spring-ioc-xml-03.xml");
        context1.refresh();

        return context;
    }

    /**
     * 读取 ioc 容器的组件
     */
    @Test
    public void getBean() {
        ApplicationContext context = createIoc();
        // 方案一: （不推荐）直接根据 beanId 获取组件，不过需要进行类型强制转换
        HappyComponent happyComponent = (HappyComponent)context.getBean("happyComponent");
        // 方案二：根据 beanId，同时指定 bean 类型
        HappyComponent happyComponent2 = context.getBean("happyComponent", HappyComponent.class);
        // 方案三：直接根据 bean 类型获取，但有以下注意事项
        // 1. 同类型的 bean 在 ioc 容器中只能存在一个
        // 2. ioc 的配置一定是类，但是可以根据接口类型(类实现的接口)获取值
        HappyComponent happyComponent3 = context.getBean(HappyComponent.class);

        happyComponent.doWork();
        happyComponent2.doWork();
        happyComponent3.doWork();
    }

    @Test
    public void beanPeriodic() {
        // 创建 ioc 容器，bean 的 init 方法将被调用
        ClassPathXmlApplicationContext context = new ClassPathXmlApplicationContext("spring-ioc-xml-04.xml");

        // 关闭 ioc 容器，bean 的 destroy 方法将被调用
        // 如果不是通过 close() 方法正常关闭 ioc 容器，那么 ioc 容器内的 beans 的 destroy 方法将不会被调用
        context.close();
    }

    @Test
    public void beanFactory() {
        // 创建 ioc 容器，bean 的 init 方法将被调用
        ClassPathXmlApplicationContext context = new ClassPathXmlApplicationContext("spring-ioc-xml-05.xml");

        // 获取所需对象的 bean
        JavaBean bean = context.getBean("bean1", JavaBean.class);
        System.out.println(bean.toString());

        // 获取工厂对象的 bean
        JavaBeanFactory beanFactory = context.getBean("&bean1", JavaBeanFactory.class);
        System.out.println(beanFactory.toString());

        // 关闭 ioc 容器，bean 的 destroy 方法将被调用
        // 如果不是通过 close() 方法正常关闭 ioc 容器，那么 ioc 容器内的 beans 的 destroy 方法将不会被调用
        context.close();
    }


}
