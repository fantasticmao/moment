package cn.fantasticmao.util.moment;

import com.beust.jcommander.JCommander;
import com.beust.jcommander.ParameterException;

import java.io.File;
import java.util.List;
import java.util.stream.Collectors;

/**
 * Main
 *
 * @author maodh
 * @since 2018/9/4
 */
public class Main {

    public static void main(String[] args) {
        Args arguments = new Args();
        try {
            JCommander jCommander = JCommander.newBuilder().addObject(arguments).build();
            jCommander.parse(args);
        } catch (ParameterException e) {
            e.usage();
            return;
        }

        final List<File> imageFileList = arguments.getPathList().stream().map(File::new).collect(Collectors.toList());
        final int subtitleHeight = arguments.getHeight();
        final String outputDir = arguments.getOutputDir();

        try {
            File finalImageFile = StitchUtil.stitchImageList(imageFileList, subtitleHeight, outputDir);
            System.out.println("well done, you can find your image here [ " + finalImageFile.getPath() + " ]");
        } catch (IllegalArgumentException e) {
            System.out.println("failed, illegal argument error suggestion: [ " + e.getMessage() + " ]");
        } catch (Exception e) {
            System.out.println("failed, report the problem mailto:maomao8017@gmail.com");
            e.printStackTrace();
        }
    }
}
