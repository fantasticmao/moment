package cn.fantasticmao.util.moment;

import com.beust.jcommander.Parameter;
import lombok.Data;

import java.io.File;
import java.util.List;

/**
 * Args 命令行输入参数
 *
 * @author maodh
 * @since 2018/9/4
 */
@Data
class Args {
    @Parameter(names = {"-p", "--path"}, description = "The path of image list that need to be stitched", required = true, order = 0)
    private List<String> pathList;
    @Parameter(names = {"-h", "--height"}, description = "The bottom subtitle height should be in unit px", order = 1)
    private Integer height = 120;
    @Parameter(names = "--out", description = "The directory to save the finished image", order = 2)
    private String outputDir = System.getProperty("user.dir") + File.separator;
}