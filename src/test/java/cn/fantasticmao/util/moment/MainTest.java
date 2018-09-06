package cn.fantasticmao.util.moment;

import org.junit.Assert;
import org.junit.Test;

import java.io.File;

/**
 * MainTest
 *
 * @author maodh
 * @since 2018/9/5
 */
public class MainTest {

    @Test
    public void main() {
        final String testImageBaseDir = this.getClass().getResource("/").getPath() + File.separator;

        Main.main(new String[]{"-p", testImageBaseDir + "1.png",
                "-p", testImageBaseDir + "2.png",
                "-p", testImageBaseDir + "3.png",
                "-p", testImageBaseDir + "4.png",
                "-p", testImageBaseDir + "5.png",
                "-h", "120",
                "--out", testImageBaseDir});

        File finalImage = new File(testImageBaseDir + "1_final.png");
        Assert.assertTrue(finalImage.exists() && finalImage.isFile());
    }
}