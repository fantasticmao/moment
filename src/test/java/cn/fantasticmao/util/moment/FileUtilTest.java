package cn.fantasticmao.util.moment;

import org.junit.Assert;
import org.junit.Test;

import java.io.File;
import java.nio.file.Paths;

/**
 * FileUtilTest
 *
 * @author maodh
 * @since 2018/9/5
 */
public class FileUtilTest {
    private final String filePath = this.getClass().getResource("/").getPath() + File.separator + "1.png";

    @Test
    public void getFileNameWithoutExtension() {
        File file = Paths.get(filePath).toFile();
        String fileNameWithoutExtension = FileUtil.getFileNameWithoutExtension(file);
        Assert.assertEquals("1", fileNameWithoutExtension);
    }

    @Test
    public void getFileExtension() {
        File file = Paths.get(filePath).toFile();
        String fileExtension = FileUtil.getFileExtension(file);
        Assert.assertEquals(".png", fileExtension);
    }

    @Test
    public void getFileExtensionWithoutDot() {
        File file = Paths.get(filePath).toFile();
        String fileExtensionWithoutDot = FileUtil.getFileExtensionWithoutDot(file);
        Assert.assertEquals("png", fileExtensionWithoutDot);
    }

    @Test
    public void compareImageExtensionAndMagicNumber() {
        File file = Paths.get(filePath).toFile();
        Assert.assertTrue(FileUtil.compareImageExtensionAndMagicNumber(file));
    }
}