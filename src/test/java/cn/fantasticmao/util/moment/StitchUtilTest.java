package cn.fantasticmao.util.moment;

import org.junit.Assert;
import org.junit.Test;

import java.io.File;
import java.io.IOException;
import java.util.Collections;
import java.util.List;
import java.util.stream.Collectors;
import java.util.stream.Stream;

/**
 * StitchUtilTest
 *
 * @author maodh
 * @since 2018/9/5
 */
public class StitchUtilTest {

    @Test
    public void test() throws IOException {
        final String testImageBaseDir = this.getClass().getResource("/").getPath() + File.separator;
        final List<File> imageFileList = Stream.of(testImageBaseDir + "1.png", testImageBaseDir + "2.png",
                testImageBaseDir + "3.png", testImageBaseDir + "4.png", testImageBaseDir + "5.png")
                .map(File::new).collect(Collectors.toList());
        final int subtitleHeight = 120;

        File file = StitchUtil.stitchImageList(imageFileList, subtitleHeight, testImageBaseDir);
        Assert.assertTrue(file.exists() && file.isFile());
    }

    @Test
    public void test1() throws IOException {
        final String testImageBaseDir = this.getClass().getResource("/").getPath() + File.separator;
        final int subtitleHeight = 120;

        try {
            StitchUtil.stitchImageList(Collections.emptyList(), subtitleHeight, testImageBaseDir);
        } catch (IllegalArgumentException e) {
            Assert.assertEquals("the image list should not be empty", e.getMessage());
        }
    }

    @Test
    public void test2() throws IOException {
        final String testImageBaseDir = this.getClass().getResource("/").getPath() + File.separator;
        final List<File> imageFileList = Stream.of(testImageBaseDir + "not_exists.png")
                .map(File::new).collect(Collectors.toList());
        final int subtitleHeight = 120;

        try {
            StitchUtil.stitchImageList(imageFileList, subtitleHeight, testImageBaseDir);
        } catch (IllegalArgumentException e) {
            Assert.assertEquals("every list should be filled with image", e.getMessage());
        }
    }

    @Test
    public void test3() throws IOException {
        final String testImageBaseDir = this.getClass().getResource("/").getPath() + File.separator;
        final List<File> imageFileList = Stream.of(testImageBaseDir + "1.png")
                .map(File::new).collect(Collectors.toList());
        final int subtitleHeight = -120;

        try {
            StitchUtil.stitchImageList(imageFileList, subtitleHeight, testImageBaseDir);
        } catch (IllegalArgumentException e) {
            Assert.assertEquals("the subtitle height should be a positive integer", e.getMessage());
        }
    }

    @Test
    public void test4() throws IOException {
        final String testImageBaseDir = this.getClass().getResource("/").getPath() + File.separator;
        final List<File> imageFileList = Stream.of(testImageBaseDir + "1.png")
                .map(File::new).collect(Collectors.toList());
        final int subtitleHeight = 120;

        try {
            StitchUtil.stitchImageList(imageFileList, subtitleHeight, "");
        } catch (IllegalArgumentException e) {
            Assert.assertEquals("the output directory path should not be empty", e.getMessage());
        }
    }

    @Test
    public void test5() throws IOException {
        final String testImageBaseDir = this.getClass().getResource("/").getPath() + File.separator;
        final List<File> imageFileList = Stream.of(testImageBaseDir + "1.png")
                .map(File::new).collect(Collectors.toList());
        final int subtitleHeight = 120;

        try {
            StitchUtil.stitchImageList(imageFileList, subtitleHeight, testImageBaseDir + "not_exists");
        } catch (IllegalArgumentException e) {
            Assert.assertEquals("there must be an output directory for finished image", e.getMessage());
        }
    }
}