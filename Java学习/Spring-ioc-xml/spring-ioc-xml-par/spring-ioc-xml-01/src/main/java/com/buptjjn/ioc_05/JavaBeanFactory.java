package com.buptjjn.ioc_05;

import org.springframework.beans.factory.BeanFactory;
import org.springframework.beans.factory.FactoryBean;

public class JavaBeanFactory implements FactoryBean<JavaBean> {

    private String name;

    private int age;

    public void setAge(int age) {
        this.age = age;
    }

    public void setName(String name) {
        this.name = name;
    }

    @Override
    // 在此处完成对所需对象的构建，可以是简单的构造方法，也可以是 setter方法，或者是其他更加复杂的创建方法
    public JavaBean getObject() throws Exception {
        JavaBean bean = new JavaBean("bean", 18);

        bean.setAge(this.age);
        bean.setName(this.name);

        return bean;
    }

    @Override
    // 返回所需对象的 class 类型
    public Class<?> getObjectType() {
        return JavaBean.class;
    }

    @Override
    // 对象是否是单例模式
    public boolean isSingleton() {
        return true;
    }

    @Override
    public String toString() {
        return "JavaBeanFactory{}";
    }
}
